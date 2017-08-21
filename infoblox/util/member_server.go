package util

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common"
)

// MemberServerListSchema - returns the schema for a list of member server structs
func MemberServerListSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "The grid primary servers for this zone.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"gridreplicate": {
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
				"enablepreferredprimaries": {
					Type:        schema.TypeBool,
					Description: "This flag represents whether the preferred_primaries field values of this member are used. Defaults to false",
					Optional:    true,
				},
				"preferredprimaries": ExternalServerListSchema(true, false),
				"stealth": {
					Type:        schema.TypeBool,
					Description: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
					Optional:    true,
				},
			},
		},
	}
}

// BuildMemberServerListFromT - Build a list of member servers out of the corresponding template structure
// Input: the list of member servers as from the template
// Output:  the corresponding []zoneauth.MemberServer list
func BuildMemberServerListFromT(members []map[string]interface{}) []common.MemberServer {
	ms := []common.MemberServer{}
	for _, item := range members {

		var server common.MemberServer
		server.Name = item["name"].(string)

		if v, ok := item["gridreplicate"]; ok {
			gr := v.(bool)
			server.GridReplicate = &gr
		}

		if v, ok := item["lead"]; ok {
			ld := v.(bool)
			server.GridReplicate = &ld
		}

		if v, ok := item["enablepreferredprimaries"]; ok {
			epp := v.(bool)
			server.EnablePreferedPrimaries = &epp
		}

		if v, ok := item["preferredprimaries"].([]map[string]interface{}); ok {
			server.PreferredPrimaries = BuildExternalServerListFromT(v)
		}

		if v, ok := item["stealth"]; ok {
			st := v.(bool)
			server.Stealth = &st
		}

		ms = append(ms, server)
	}
	return ms
}

// BuildMemberServerListFromIBX - builds a list of maps that match MemberServer data model
// out of the corresponding IBX structure.
func BuildMemberServerListFromIBX(IBXmembersList []common.MemberServer) []map[string]interface{} {

	memberServerList := make([]map[string]interface{}, 0)

	for _, IBXmember := range IBXmembersList {
		member := make(map[string]interface{})

		if IBXmember.GridReplicate != nil {
			f := IBXmember.GridReplicate
			member["gridreplicate"] = *f
		}

		if IBXmember.Lead != nil {
			member["lead"] = *IBXmember.Lead
		}

		if IBXmember.Name != "" {
			member["name"] = IBXmember.Name
		}

		if IBXmember.EnablePreferedPrimaries != nil {
			member["enablepreferredprimaries"] = *IBXmember.EnablePreferedPrimaries
		}

		if IBXmember.PreferredPrimaries != nil {
			member["preferredprimaries"] = BuildExternalServersListFromIBX(IBXmember.PreferredPrimaries)
		}

		if IBXmember.Stealth != nil {
			member["stealth"] = *IBXmember.Stealth
		}

		memberServerList = append(memberServerList, member)
	}

	return memberServerList
}
