package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceZoneDelegated() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneDelegatedCreate,
		Read:   resourceZoneDelegatedRead,
		Update: resourceZoneDelegateUpdate,
		Delete: resourceZoneDelegatedDelete,
		Schema: map[string]*schema.Schema{
			"Reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"view": {
				Type:        schema.TypeString,
				Description: "The name of the DNS view in which the zone resides",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment for the zone; maximum 256 characters",
				Optional:    true,
			},
			"delegateto": {
				Type:        schema.TypeList,
				Description: "A list of name servers the zone is delegated to",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the delegation server",
						},
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ip address of the delegation server",
						},
						"stealth": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Set this flag to hide the NS record for the primary name server from DNS queries",
						},
						"tsigkey": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A generated TSIG key.",
						},
						"tsigkeyalg": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The TSIG key algorithm",
							ValidateFunc: util.ValidateTSIGAlgorithm,
						},
						"tsigkeyname": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The TSIG key name",
						},
						"usetsigkeyname": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use flag for: tsigkeyname",
						},
					},
				},
			},
			"delegatedttl": {
				Type:        schema.TypeInt,
				Description: "a TTL for the delegated zone",
				Optional:    true,
				ForceNew:    false,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Is the zone disabled",
				Optional:    true,
				ForceNew:    false,
			},
			"fqdn": {
				Type:        schema.TypeString,
				Description: "The FQDN for the zone that is being delegated",
				Required:    true,
				ForceNew:    true,
			},
			"locked": {
				Type:        schema.TypeBool,
				Description: "Is the record locked to prevent changes",
				Optional:    true,
				ForceNew:    false,
			},
			"usedelegatedttl": {
				Type:        schema.TypeBool,
				Description: "Should we use the deletated ttl",
				Optional:    true,
			},
			"zoneformat": {
				Type:         schema.TypeString,
				Description:  "Format of the zone, default is FORWARD",
				Optional:     true,
				Default:      "FORWARD",
				ValidateFunc: util.ValidateZoneFormat,
			},
			"ns_group": {
				Type:        schema.TypeString,
				Description: "NameServer group for this zone",
				Optional:    true,
			},
		},
	}
}

func resourceZoneDelegatedCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var createZoneDelegated zonedelegated.ZoneDelegated

	if v, ok := d.GetOk("fqdn"); ok {
		createZoneDelegated.Fqdn = v.(string)
	}

	if v, ok := d.GetOk("comment"); ok {
		createZoneDelegated.Comment = v.(string)
	}

	if v, ok := d.GetOk("view"); ok {
		createZoneDelegated.View = v.(string)
	}

	if v, ok := d.GetOk("disable"); ok {
		delegationDisable := v.(bool)
		createZoneDelegated.Disable = &delegationDisable
	}

	if v, ok := d.GetOk("delegateto"); ok {
		delegatedServers := []map[string]interface{}{}
		for _, delegatedServer := range v.([]interface{}) {
			delegatedServers = append(delegatedServers, delegatedServer.(map[string]interface{}))
		}
		createZoneDelegated.DelegateTo = util.BuildExternalServerListFromT(delegatedServers)
	}

	if v, ok := d.GetOk("delegatedttl"); ok {
		createZoneDelegated.DelegatedTTL = uint(v.(int))
	}

	if v, ok := d.GetOk("usedelegatedttl"); ok {
		useTTL := v.(bool)
		createZoneDelegated.UseDelegatedTTL = &useTTL
	}

	if v, ok := d.GetOk("locked"); ok {
		locked := v.(bool)
		createZoneDelegated.Locked = &locked
	}

	if v, ok := d.GetOk("zoneformat"); ok {
		createZoneDelegated.ZoneFormat = v.(string)
	}

	if v, ok := d.GetOk("ns_group"); ok {
		createZoneDelegated.NsGroup = v.(string)
	}

	createZoneDeletagedAPI := zonedelegated.NewCreate(createZoneDelegated)
	errCreate := infobloxClient.Do(createZoneDeletagedAPI)
	if errCreate != nil {
		return fmt.Errorf("Error creating Zone Delegated %s", errCreate.Error())
	}
	if createZoneDeletagedAPI.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Error creating Zone Delegated %s", *createZoneDeletagedAPI.ResponseObject().(*string))
	}
	d.SetId(*createZoneDeletagedAPI.ResponseObject().(*string))
	return resourceZoneDelegatedRead(d, m)

}

func resourceZoneDelegatedRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var readZoneDelegated zonedelegated.ZoneDelegated
	returnFields := []string{"address", "comment", "fqdn", "disable", "zone_format", "delegate_to", "delegated_ttl", "locked", "use_delegated_ttl", "ns_group"}
	readAPI := zonedelegated.NewGet(d.Id(), returnFields)
	readErr := infobloxClient.Do(readAPI)
	if readErr != nil {
		return fmt.Errorf("Could not read the resource %s", readErr.Error())
	}

	if readAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Could not read the resource %s", *readAPI.ResponseObject().(*string))
	}
	readZoneDelegated = *readAPI.ResponseObject().(*zonedelegated.ZoneDelegated)
	d.SetId(readZoneDelegated.Ref)
	d.Set("comment", readZoneDelegated.Comment)
	d.Set("view", readZoneDelegated.View)
	d.Set("delegateto", readZoneDelegated.DelegateTo)
	d.Set("delegatedttl", readZoneDelegated.DelegatedTTL)
	d.Set("disable", readZoneDelegated.Disable)
	d.Set("fqdn", readZoneDelegated.Fqdn)
	d.Set("locked", readZoneDelegated.Locked)
	d.Set("usedelegatedttl", readZoneDelegated.UseDelegatedTTL)
	d.Set("zoneformat", readZoneDelegated.ZoneFormat)
	d.Set("ns_group", readZoneDelegated.NsGroup)
	return nil
}

func resourceZoneDelegateUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var updateZoneDelegated zonedelegated.ZoneDelegated
	var hasChange bool
	if d.HasChange("comment") {
		_, newComment := d.GetChange("comment")
		updateZoneDelegated.Comment = newComment.(string)
		hasChange = true
	}
	if d.HasChange("view") {
		_, newView := d.GetChange("view")
		updateZoneDelegated.View = newView.(string)
		hasChange = true
	}
	if d.HasChange("delegateto") {
		_, newDelegateto := d.GetChange("delegateto")
		delegatedServers := []map[string]interface{}{}
		for _, delegatedServer := range newDelegateto.([]interface{}) {
			delegatedServers = append(delegatedServers, delegatedServer.(map[string]interface{}))
		}
		updateZoneDelegated.DelegateTo = util.BuildExternalServerListFromT(delegatedServers)
		hasChange = true
	}

	if d.HasChange("delegatettl") {
		_, newDelegateTTL := d.GetChange("delegatettl")
		updateZoneDelegated.DelegatedTTL = newDelegateTTL.(uint)
		hasChange = true
	}

	if d.HasChange("disable") {
		var disabled bool
		_, newDisable := d.GetChange("disable")
		disabled = newDisable.(bool)
		updateZoneDelegated.Disable = &disabled
		hasChange = true
	}

	if d.HasChange("locked") {
		var isLocked bool
		_, newLocked := d.GetChange("locked")
		isLocked = newLocked.(bool)
		updateZoneDelegated.Locked = &isLocked
		hasChange = true
	}

	if d.HasChange("usedelegatedttl") {
		var useDelegated bool
		_, newUseDelegatedTTL := d.GetChange("usedelegatedttl")
		useDelegated = newUseDelegatedTTL.(bool)
		updateZoneDelegated.UseDelegatedTTL = &useDelegated
		hasChange = true
	}

	if d.HasChange("zoneformat") {
		_, newZoneFormat := d.GetChange("zoneformat")
		updateZoneDelegated.ZoneFormat = newZoneFormat.(string)
		hasChange = true
	}

	if d.HasChange("ns_group") {
		_, newNsGroup := d.GetChange("ns_group")
		updateZoneDelegated.NsGroup = newNsGroup.(string)
		hasChange = true
	}

	if hasChange {
		updateZoneDelegatedAPI := zonedelegated.NewUpdate(d.Id(), updateZoneDelegated)
		updateZoneErr := infobloxClient.Do(updateZoneDelegatedAPI)
		if updateZoneErr != nil {
			return fmt.Errorf("Could not update the zone : %s", updateZoneErr.Error())
		}
		if updateZoneDelegatedAPI.StatusCode() != http.StatusOK {
			return fmt.Errorf("Infoblox Zone Auth Update return code != 200")
		}
		return resourceZoneDelegatedRead(d, m)
	}

	return nil

}

func resourceZoneDelegatedDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	deleteZoneDelegatedAPI := zonedelegated.NewDelete(d.Id())
	deleteErr := infobloxClient.Do(deleteZoneDelegatedAPI)
	if deleteErr != nil {
		return fmt.Errorf("Could not delete the record: %s ", deleteErr.Error())
	}
	d.SetId("")
	return nil
}
