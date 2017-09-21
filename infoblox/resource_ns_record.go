package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSRecordCreate,
		Read:   resourceNSRecordRead,
		Update: resourceNSRecordUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"addresses": {
				Type:        schema.TypeSet,
				Description: "The list of zone name servers",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The address of the Zone Name Server",
							Required:    true,
						},
						"auto_create_ptr": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Flag to indicate if PTR records need to be auto created",
							Optional:    true,
						},
					},
				},
			},
			"creator": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC",
			},
			"dns_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "The name for the NS record in punycode format. Values with leading or trailing white space are not valid for this field. Cannot be written nor updated.",
			},
			"name": {
				Type:         schema.TypeString,
				Description:  "The name of the ns record where the record should reside. Cannot be updated",
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"ms_delegation_name": {
				Type:        schema.TypeString,
				Description: "The MS delegation point name",
				Optional:    true,
			},
			"nameserver": {
				Type:         schema.TypeString,
				Description:  "The FQDN of the name server",
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"view": {
				Type:        schema.TypeString,
				Description: "The name of the DNS view in which the record resides",
				Optional:    true,
				Computed:    true,
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS Zone for the record",
			},
		},
	}
}

func resourceNSRecordCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.RecordNSObj, resourceNSRecord(), d, m)
}

func resourceNSRecordRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceNSRecord(), d, m)
}

func resourceNSRecordUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceNSRecord(), d, m)
}
