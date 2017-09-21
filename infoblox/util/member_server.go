package util

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// MemberServerListSchema - returns the schema for a list of member server structs
func MemberServerListSchema(optional, required bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "The grid primary servers for this zone.",
		Optional:    optional,
		Required:    required,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"grid_replicate": {
					Type:        schema.TypeBool,
					Description: "The flag represents DNS zone transfers if set to True, and ID Grid Replication if set to False. This flag is ignored if the struct is specified as part of a stub zone or if it is set as grid_member in an authoritative zone.",
					Optional:    true,
				},
				"lead": {
					Type:        schema.TypeBool,
					Description: "This flag controls whether the Grid lead secondary server performs zone transfers to non lead secondaries. This flag is ignored if the struct is specified as grid_member in an authoritative zone.",
					Optional:    true,
				},
				"name": {
					Type:        schema.TypeString,
					Description: "The grid member name.",
					Required:    true,
				},
				"enable_preferred_primaries": {
					Type:        schema.TypeBool,
					Description: "This flag represents whether the preferred_primaries field values of this member are used. Defaults to false",
					Optional:    true,
				},
				"preferred_primaries": ExternalServerListSchema(true, false),
				"stealth": {
					Type:        schema.TypeBool,
					Description: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
					Optional:    true,
				},
			},
		},
	}
}
