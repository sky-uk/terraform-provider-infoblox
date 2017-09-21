package util

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// TODO - methods TBD

// ForwardingMemberServerListSchema - returns a list of Forwarding Member Servers
func ForwardingMemberServerListSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "The primary preference list with Grid member names and/or External Server structs for this member.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"forward_to": ExternalServerListSchema(true, false),
				"forwarders_only": {
					Type:        schema.TypeBool,
					Description: "Determines if the appliance sends queries to forwarders only and not to other internal or Internet root servers",
					Optional:    true,
				},
				"name": {
					Type:        schema.TypeString,
					Description: "The name of this Grid member in FQDN format.The field is required on creation.",
					Required:    true,
				},
				"use_override_forwarders": {
					Type:        schema.TypeBool,
					Description: "Use flag for: forward_to.The default value is False.",
					Optional:    true,
				},
			},
		},
	}
}
