package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceSRVRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceSRVRecordCreate,
		Read:   resourceSRVRecordRead,
		Update: resourceSRVRecordUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for a SRV record in FQDN format",
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"target": {
				Type:     schema.TypeString,
				Required: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "The name of the DNS View in which the record resides",
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The Time To Live assigned to CNAME",
			},
			"use_ttl": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the record; maximum 256 characters",
			},
		},
	}
}

func resourceSRVRecordCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.RecordSRVObj, resourceSRVRecord(), d, m)
}

func resourceSRVRecordRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceSRVRecord(), d, m)
}

func resourceSRVRecordUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceSRVRecord(), d, m)
}
