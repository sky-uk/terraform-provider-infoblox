package util

import (
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildMemberServerListFromT(t *testing.T) {
	preferredPrimaries := make([]map[string]interface{}, 0)
	preferredPrimary := make(map[string]interface{})
	preferredPrimary["name"] = "foo"
	preferredPrimaries = append(preferredPrimaries, preferredPrimary)

	var IBXprefPrim common.ExternalServer
	IBXprefPrim.Name = "foo"
	IBXprefPrims := make([]common.ExternalServer, 0)
	IBXprefPrims = append(IBXprefPrims, IBXprefPrim)
	IBXMemberList := make([]common.MemberServer, 0)
	IBXmember := common.MemberServer{
		Name:               "foo",
		PreferredPrimaries: IBXprefPrims,
	}
	IBXMemberList = append(IBXMemberList, IBXmember)

	memberListMap := make([]map[string]interface{}, 1)
	member := make(map[string]interface{})
	member["name"] = "foo"
	member["preferredprimaries"] = preferredPrimaries
	memberListMap[0] = member

	memberServerList := BuildMemberServerListFromT(memberListMap)
	assert.Equal(t, memberServerList, IBXMemberList)
}
func TestBuildMemberServerListFromIBX(t *testing.T) {
	IBXmemberList := make([]common.MemberServer, 0)

	b := true
	IBXmemberItem := common.MemberServer{
		GridReplicate: &b,
		Lead:          &b,
		Name:          "foo",
		EnablePreferedPrimaries: &b,
		PreferredPrimaries: []common.ExternalServer{
			common.ExternalServer{
				Address:        "10.10.10.10",
				Name:           "bar",
				Stealth:        &b,
				TsigKey:        "baz",
				TsigKeyAlg:     "HMAC",
				TsigKeyName:    "foobar",
				UseTsigKeyName: &b,
			}},
		Stealth: &b,
	}

	IBXmemberList = append(IBXmemberList, IBXmemberItem)
	IBXmemberList = append(IBXmemberList, IBXmemberItem)

	templateMapList := BuildMemberServerListFromIBX(IBXmemberList)

	assert.Equal(t, 2, len(templateMapList))
	assert.Equal(t, b, templateMapList[0]["gridreplicate"])
	assert.Equal(t, b, templateMapList[0]["lead"])
	assert.Equal(t, "foo", templateMapList[0]["name"])
	assert.Equal(t, "foo", templateMapList[1]["name"])
	assert.Equal(t, b, templateMapList[0]["enablepreferredprimaries"])
	preferredPrimaries := templateMapList[0]["preferredprimaries"].([]map[string]interface{})
	assert.Equal(t, "10.10.10.10", preferredPrimaries[0]["address"])
	assert.Equal(t, "bar", preferredPrimaries[0]["name"])
	assert.Equal(t, b, preferredPrimaries[0]["stealth"])
	assert.Equal(t, "baz", preferredPrimaries[0]["tsig_key"])
	assert.Equal(t, "HMAC", preferredPrimaries[0]["tsig_key_alg"])
	assert.Equal(t, "foobar", preferredPrimaries[0]["tsig_key_name"])
	assert.Equal(t, b, preferredPrimaries[0]["use_tsig_key_name"])
}
