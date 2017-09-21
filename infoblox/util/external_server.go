package util

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// ExternalServerListSchema - returns the schema for a list of external servers
func ExternalServerListSchema(optional, required bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "The primary preference list with Grid member names and/or External Server structs for this member.",
		Optional:    optional,
		Required:    required,
		Elem:        externalServerSchema(),
	}
}

// ExternalServerSetSchema - returns the schema for a set of external servers
func ExternalServerSetSchema(optional, required bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Description: "The primary preference set with Grid member names and/or External Server structs for this member.",
		Optional:    optional,
		Required:    required,
		Elem:        externalServerSchema(),
	}
}

// externalServerSchema - returns an external server resource
func externalServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Description: "The IPv4 Address or IPv6 Address of the server.",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "A resolvable domain name for the external DNS server.",
				Required:    true,
			},
			"shared_with_ms_parent_delegation": {
				Type:        schema.TypeBool,
				Description: "This flag represents whether the name server is shared with the parent Microsoft primary zone’s delegation server.",
				Optional:    true,
				Computed:    true,
			},
			"stealth": {
				Type:        schema.TypeBool,
				Description: "Set this flag to hide the NS record for the primary name server from DNS queries.",
				Optional:    true,
			},
			"tsig_key": {
				Type:        schema.TypeString,
				Description: "A generated TSIG key. Values with leading or trailing whitespace are not valid for this field.",
				Optional:    true,
			},
			"tsig_key_alg": {
				Type:        schema.TypeString,
				Description: "The TSIG key algorithm. Valid values: HMAC-MD5 or HMAC-SHA256. The default value is HMAC-MD5.",
				Optional:    true,
				Default:     "HMAC-MD5",
			},
			"tsig_key_name": {
				Type:        schema.TypeString,
				Description: "The TSIG key name.",
				Optional:    true,
			},
			"use_tsig_key_name": {
				Type:        schema.TypeBool,
				Description: "Use flag for: tsig_key_name",
				Optional:    true,
			},
		},
	}
}
