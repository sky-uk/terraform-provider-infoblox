package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dhcp_range"
	"log"
	"net/http"
)

func resourceDHCPRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPRangeCreate,
		Read:   resourceDHCPRangeRead,
		Delete: resourceDHCPRangeDelete,
		Update: resourceDHCPRangeUpdate,

		Schema: map[string]*schema.Schema{
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique reference to Infoblox Network resource",
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"network_view": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "default",
			},
			"start": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"end": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"member": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    false,
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
				ForceNew:    false,
				Description: "Restarts any services if required by this change. Default: true.",
				Default:     false,
			},
			"server_association": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
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

	if v, ok := d.GetOk("name"); ok {
		rangeCreate.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok {
		rangeCreate.Comment = v.(string)
	}

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
		rangeRestart := v.(bool)
		rangeCreate.Restart = &rangeRestart
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
	fields := []string{"name", "comment", "end_addr", "start_addr", "network", "network_view", "member", "server_association_type"}
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
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	return nil
}

//resourceDHCPRangeUpdate - update the DHCPRange object
func resourceDHCPRangeUpdate(d *schema.ResourceData, m interface{}) error {
	var rangeUpdate dhcprange.DHCPRange
	var hasChanges bool
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	fields := []string{"name", "comment", "end_addr", "start_addr", "network", "network_view", "member", "server_association_type"}
	rangeUpdateAPI := dhcprange.NewGetDHCPRangeAPI(d.Id(), fields)
	getErr := infobloxClient.Do(rangeUpdateAPI)
	if getErr != nil {
		return fmt.Errorf("Could not read resource %s", getErr)
	}
	rangeUpdate = rangeUpdateAPI.GetResponse()
	//rangeUpdate.Ref = d.Id()
	if d.HasChange("name") {
		hasChanges = true
		_, newName := d.GetChange("name")
		rangeUpdate.Name = newName.(string)
	}
	if d.HasChange("comment") {
		hasChanges = true
		_, newComment := d.GetChange("comment")
		rangeUpdate.Comment = newComment.(string)
	}
	if d.HasChange("network") {
		hasChanges = true
		_, newNetwork := d.GetChange("network")
		rangeUpdate.Network = newNetwork.(string)
	}
	if d.HasChange("network_view") {
		hasChanges = true
		_, newNetworkView := d.GetChange("network_view")
		rangeUpdate.NetworkView = newNetworkView.(string)
	}

	if d.HasChange("start") {
		hasChanges = true
		_, newStart := d.GetChange("start")
		rangeUpdate.Start = newStart.(string)
	}

	if d.HasChange("end") {
		hasChanges = true
		_, newEnd := d.GetChange("end")
		rangeUpdate.End = newEnd.(string)
	}

	if d.HasChange("member") {
		hasChanges = true
		_, newMember := d.GetChange("member")
		rangeUpdate.Member = newMember.(dhcprange.Member)

	}

	if d.HasChange("restart") {
		hasChanges = true
		_, newRestart := d.GetChange("restart")
		restart := newRestart.(bool)
		rangeUpdate.Restart = &restart
	}

	if d.HasChange("server_association") {
		hasChanges = true
		_, newServerAssociation := d.GetChange("server_association")
		rangeUpdate.ServerAssociation = newServerAssociation.(string)
	}
	if hasChanges {
		updateRangeAPI := dhcprange.NewUpdateDHCPRange(rangeUpdate)
		updateErr := infobloxClient.Do(updateRangeAPI)
		if updateErr != nil {

			return fmt.Errorf("cound not update the dhcprange , status code:  %d", updateRangeAPI.StatusCode())
		}
		if updateRangeAPI.StatusCode() != 200 {
			log.Println(rangeUpdate)
			log.Println("Endpoint used: " + updateRangeAPI.Endpoint())
			log.Println("Response :" + updateRangeAPI.GetResponse())
			return fmt.Errorf("cound not update the dhcprange , status code:  %d", updateRangeAPI.StatusCode())
		}
		d.SetId(updateRangeAPI.GetResponse())
		return resourceDHCPRangeRead(d, m)
	}
	return nil
}

// buildMemberObject - This is to avoid having to repeat the code every time I need to read this field
func buildMemberObject(memberSet *schema.Set) dhcprange.Member {

	member := dhcprange.Member{ElementType: "dhcpmember"}

	for _, object := range memberSet.List() {
		memberObject := object.(map[string]interface{})
		if address, ok := memberObject["ipv4_addr"].(string); ok {
			member.IPv4Address = address
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
	r["ipv4_addr"] = member.IPv4Address
	r["name"] = member.Name
	result = append(result, r)
	return result
}
