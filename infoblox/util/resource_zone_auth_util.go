package util

import (
	"github.com/sky-uk/skyinfoblox/api/common"
	"log"
)

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
			log.Printf("----------------> ITEMS:\n%+v\n", v)
			if len(v) > 0 {
				servers := make([]map[string]interface{}, 0)
				for _, server := range v {
					servers = append(servers, server)
				}
				server.PreferredPrimaries = BuildExternalServerListFromT(servers)
			}
		}

		if v, ok := item["stealth"]; ok {
			st := v.(bool)
			server.Stealth = &st
		}

		ms = append(ms, server)
	}
	return ms
}

// BuildExternalServerListFromT - Builds a list of external servers given the corresponding list of
// items from state
func BuildExternalServerListFromT(extServerListFromT []map[string]interface{}) []common.ExternalServer {

	es := []common.ExternalServer{}
	for _, item := range extServerListFromT {
		var extServer common.ExternalServer

		if v, ok := item["address"]; ok {
			extServer.Address = v.(string)
		}

		if v, ok := item["name"]; ok {
			extServer.Name = v.(string)
		}

		if v, ok := item["stealth"]; ok {
			b := v.(bool)
			extServer.Stealth = &b
		}

		if v, ok := item["tsigkey"]; ok {
			extServer.TsigKey = v.(string)
		}

		if v, ok := item["tsigkeyAlg"]; ok {
			extServer.TsigKeyAlg = v.(string)
		}

		if v, ok := item["tsigkeyName"]; ok {
			extServer.TsigKeyName = v.(string)
		}

		if v, ok := item["usetsigkeyname"]; ok {
			b := v.(bool)
			extServer.UseTsigKeyName = &b
		}

		es = append(es, extServer)
	}
	return es
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

// BuildExternalServersListFromIBX - builds a list of external servers for terraform
// given the corresponding struct from IBX
func BuildExternalServersListFromIBX(IBXExtServersList []common.ExternalServer) []map[string]interface{} {
	es := make([]map[string]interface{}, 0)
	for _, IBXExtServer := range IBXExtServersList {
		server := make(map[string]interface{})

		if IBXExtServer.Address != "" {
			server["address"] = IBXExtServer.Address
		}

		if IBXExtServer.Name != "" {
			server["name"] = IBXExtServer.Name
		}

		if IBXExtServer.Stealth != nil {
			server["stealth"] = *IBXExtServer.Stealth
		}

		if IBXExtServer.TsigKey != "" {
			server["tsigkey"] = IBXExtServer.TsigKey
		}

		if IBXExtServer.TsigKeyAlg != "" {
			server["tsigkeyalg"] = IBXExtServer.TsigKeyAlg
		}

		if IBXExtServer.TsigKeyName != "" {
			server["tsigkeyname"] = IBXExtServer.TsigKeyName
		}

		if IBXExtServer.UseTsigKeyName != nil {
			server["usetsigkeyname"] = *IBXExtServer.UseTsigKeyName
		}

		es = append(es, server)
	}

	return es
}
