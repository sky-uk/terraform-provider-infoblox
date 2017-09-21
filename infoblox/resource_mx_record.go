package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"time"
)

func resourceMxRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceMxRecordCreate,
		Read:   resourceMxRecordRead,
		Update: resourceMxRecordUpdate,
		Delete: DeleteResource,

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(2 * time.Minute),
			Create: schema.DefaultTimeout(2 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "Name of the zone the MX record refers to",
			},
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "A comment on the record",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is the record disabled",
			},
			"ddns_principal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The GSS-TSIG principal that owns this record",
			},
			"ddns_protected": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the DDNS updates for this record are allowed or not",
			},
			"mail_exchanger": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "Mail exchanger name in FQDN format. This value can be in unicode format",
			},
			"preference": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format",
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The Time To Live value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid ",
			},
			"use_ttl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use flag for: ttl",
			},
			"view": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "The name of the DNS view in which the record resides.",
			},
		},
	}

}

func resourceMxRecordCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.RecordMXObj, resourceMxRecord(), d, m)
}

func resourceMxRecordRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceMxRecord(), d, m)
}

func resourceMxRecordUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceMxRecord(), d, m)
}
