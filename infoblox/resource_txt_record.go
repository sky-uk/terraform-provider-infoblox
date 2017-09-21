package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceTXTRecordCreate,
		Read:   resourceTXTRecordRead,
		Update: resourceTXTRecordUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"text": {
				Type:     schema.TypeString,
				Required: true,
			},
			"view": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"use_ttl": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceTXTRecordCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.RecordTXTObj, resourceTXTRecord(), d, m)
}

func resourceTXTRecordRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceTXTRecord(), d, m)
}

func resourceTXTRecordUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceTXTRecord(), d, m)
}
