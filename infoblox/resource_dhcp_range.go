package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dhcp_range"
	"net/http"
)

func resourceDHCPRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPRangeCreate,
		Read:   resourceDHCPRangeRead,
		Delete: resourceDHCPRangeDelete,

		Schema: map[string]*schema.Schema{
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique reference to Infoblox Network resource",
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_view": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "default",
			},
			"start": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"end": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"member": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Infoblox DHCP member that serves this range",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_addr": &schema.Schema{
							Type:        schema.TypeString,
							Description: "DHCP IPv4 Address",
							Required:    true,
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "DHCP Member server FQDN",
							Required:    true,
						},
					},
				},
			},
			"restart": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Restarts any services if required by this change. Default: true.",
				Default:     true,
			},
			"server_association": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Must be set to 'MEMBER' if member is specified",
				Default:     "NONE",
			},
		},
	}
}

// resourceDHCPRangeCreate  - Creates a new dhcp range resource
func resourceDHCPRangeCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var rangeCreate dhcprange.DHCPRange

	if v, ok := d.GetOk("network"); ok {
		rangeCreate.Network = v.(string)
	}
	if v, ok := d.GetOk("network_view"); ok {
		rangeCreate.NetworkView = v.(string)
	}
	if v, ok := d.GetOk("start"); ok {
		rangeCreate.Start = v.(string)
	}
	if v, ok := d.GetOk("end"); ok {
		rangeCreate.End = v.(string)
	}

	if v, ok := d.GetOk("member"); ok {
		if member, ok := v.(*schema.Set); ok {
			rangeCreate.Member = buildMemberObject(member)
			rangeCreate.ServerAssociation = "MEMBER"
		}
	}

	if v, ok := d.GetOk("restart"); ok {
		rangeCreate.Restart = v.(bool)
	}

	createDHCPRangeAPI := dhcprange.NewCreateDHCPRange(rangeCreate)
	err := infobloxClient.Do(createDHCPRangeAPI)
	if err != nil {
		return fmt.Errorf("Error during the DHCP Range creation request: %s", err)
	}

	if createDHCPRangeAPI.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Error creating the DHCP Range:\n %s", createDHCPRangeAPI.GetResponse())
	}
	d.SetId(createDHCPRangeAPI.GetResponse())
	return resourceDHCPRangeRead(d, m)
}

// resourceDHCPRangeDelete  - Delete a network resource
func resourceDHCPRangeDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	deleteRequest := dhcprange.NewDeleteDHCPRange(d.Id())
	deleteErr := infobloxClient.Do(deleteRequest)
	if deleteErr != nil {
		return fmt.Errorf("Cound not delete the DHCP range %s", deleteErr)
	}
	if deleteRequest.StatusCode() != http.StatusOK {
		return fmt.Errorf("Error Deleting the DHCP range: %s ", deleteRequest.GetResponse())
	}
	d.SetId("")
	return nil
}

// resourceDHCPRangeRead - Reads the resource
func resourceDHCPRangeRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	fields := []string{"end_addr", "start_addr", "network", "network_view", "member", "server_association_type"}
	getDHCPRangeRequest := dhcprange.NewGetDHCPRangeAPI(d.Id(), fields)
	getErr := infobloxClient.Do(getDHCPRangeRequest)
	if getErr != nil {
		return fmt.Errorf("Could not read resource %s", getErr)
	}

	if getDHCPRangeRequest.StatusCode() != http.StatusOK {
		return fmt.Errorf("HTTP error reading the resource:\n%s", *getDHCPRangeRequest.ResponseObject().(*string))
	}
	response := getDHCPRangeRequest.GetResponse()
	d.Set("end_addr", response.End)
	d.Set("start_addr", response.Start)
	d.Set("network", response.Network)
	d.Set("network_view", response.NetworkView)
	d.Set("server_association", response.ServerAssociation)
	d.Set("member", flattenMember(&response.Member))
	d.Set("ref", response.Ref)

	return nil
}

// buildMemberObject - This is to avoid having to repeat the code every time I need to read this field
func buildMemberObject(memberSet *schema.Set) dhcprange.Member {

	member := dhcprange.Member{InternalType: "dhcpmember"}

	for _, object := range memberSet.List() {
		memberObject := object.(map[string]interface{})
		if address, ok := memberObject["ipv4_addr"].(string); ok {
			member.Address = address
		}
		if memberFQDN, ok := memberObject["name"].(string); ok {
			member.Name = memberFQDN
		}
	}
	return member
}

// Flattens member object into a map[string]interface{}
func flattenMember(member *dhcprange.Member) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, 1)
	if member == nil {
		return nil
	}
	r := make(map[string]interface{})
	r["ipv4_addr"] = member.Address
	r["name"] = member.Name
	result = append(result, r)
	return result
}
