package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceARecordCreate,
		Read:   resourceARecordRead,
		Update: resourceARecordUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateMaxLength(256),
				Description:  "Comment for the record; maximum 256 characters", // TODO add validation function
			},
			"creation_time": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "The time of the record creation in Epoch seconds format.",
			},
			"creator": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The record creator.Valid values:DYNAMIC,STATIC,SYSTEM.Defaults to STATIC",
			},
			"ddns_protected": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the DDNS updates for this record are allowed or not",
			},
			"ddns_principal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The GSS-TSIG principal that owns this record",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the record is disabled or not. False means that the record is enabled.",
			},
			"dns_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name for an A record in punycode format. Values with leading or trailing white space are not valid for this field. Cannot be written nor updated.",
			},
			"forbid_reclamation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the reclamation is allowed for the record or not.",
			},
			"ipv4addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address for hostname",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "Name for A record in FQDN format. This value can be in unicode format. Values with leading or trailing white space are not valid for this field.",
			},
			"reclaimable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Determines if the record is reclaimable or not. Cannot be updated/written",
			},
			"shared_record_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record. Cannot be updated/written",
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "TTL in seconds for host record",
			},
			"use_ttl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Use flag for: ttl",
			},
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the DNS view in which the record resides. Example: “external”.",
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS Zone for the record",
			},
		},
	}
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.RecordAObj, resourceARecord(), d, m)
}

func resourceARecordRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceARecord(), d, m)
}

func resourceARecordUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceARecord(), d, m)
}
