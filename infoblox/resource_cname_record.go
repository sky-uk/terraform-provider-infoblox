package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceCNAMERecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceCNAMECreate,
		Read:   resourceCNAMERead,
		Update: resourceCNAMEUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"canonical": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Canonical name in FQDN format",
			},
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateMaxLength(256),
				Description:  "Comment for the record; maximum 256 characters",
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
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the record is disabled or not. False means that the record is enabled.",
			},
			"dns_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name for the CNAME record in punycode format. Values with leading or trailing white space are not valid for this field. Cannot be written nor updated.",
			},
			"forbid_reclamation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the reclamation is allowed for the record or not.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for a CNAME record in FQDN format",
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
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "The name of the DNS View in which the record resides",
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The Time To Live assigned to CNAME",
			},
			"use_ttl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Use flag for: ttl",
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS Zone for the record",
			},
		},
	}
}

func resourceCNAMECreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.RecordCnameObj, resourceCNAMERecord(), d, m)
}

func resourceCNAMERead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceCNAMERecord(), d, m)
}

func resourceCNAMEUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceCNAMERecord(), d, m)
}
