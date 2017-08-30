package util

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common"
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

// BuildForwardingMemberServerListFromT - Builds a list of forwarding member servers given the corresponding list of
// items from state
func BuildForwardingMemberServerListFromT(serverListFromT []map[string]interface{}) []common.ForwardingMemberServer {

	serverList := []common.ForwardingMemberServer{}
	for _, item := range serverListFromT {
		var server common.ForwardingMemberServer
		falseFlag := false

		if v, ok := item["name"]; ok {
			server.Name = v.(string)
		}

		if v, ok := item["forward_to"]; ok {
			serverList := GetMapList(v.([]interface{}))
			server.ForwardTo = BuildExternalServerListFromT(serverList)
		}

		if v, ok := item["forwarders_only"]; ok {
			flag := v.(bool)
			server.ForwardersOnly = &flag
		} else {
			server.ForwardersOnly = &falseFlag
		}

		if v, ok := item["use_override_forwarders"]; ok {
			flag := v.(bool)
			server.UseOverrideForwarders = &flag
		} else {
			server.UseOverrideForwarders = &falseFlag
		}

		serverList = append(serverList, server)
	}
	return serverList
}

// BuildForwardingMemberServerListFromIBX -  builds a list of forwarding member servers for terraform given
// the corresponding struct from IBX
func BuildForwardingMemberServerListFromIBX(ibxFwdMemberServerList []common.ForwardingMemberServer) []map[string]interface{} {

	forwardMemberServers := make([]map[string]interface{}, 0)
	for _, fwdMemberServer := range ibxFwdMemberServerList {
		server := make(map[string]interface{})

		if fwdMemberServer.Name != "" {
			server["name"] = fwdMemberServer.Name
		}
		if fwdMemberServer.ForwardTo != nil {
			server["forward_to"] = BuildExternalServersListFromIBX(fwdMemberServer.ForwardTo)
		}
		if fwdMemberServer.ForwardersOnly != nil {
			server["forwarders_only"] = *fwdMemberServer.ForwardersOnly
		}
		if fwdMemberServer.UseOverrideForwarders != nil {
			server["use_override_forwarders"] = *fwdMemberServer.UseOverrideForwarders
		}
		forwardMemberServers = append(forwardMemberServers, server)
	}
	return forwardMemberServers
}
