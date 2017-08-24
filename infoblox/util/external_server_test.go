package util

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildExternalServerListFromT(t *testing.T) {

	IBXExternalServersList := make([]common.ExternalServer, 1)
	b := true
	IBXExternalServer := common.ExternalServer{
		Address: "10.10.10.10",
		Name:    "Foo",
		Stealth: &b,
	}
	IBXExternalServersList[0] = IBXExternalServer

	externalServers := make([]map[string]interface{}, 1)
	member := make(map[string]interface{})
	member["name"] = "Foo"
	member["address"] = "10.10.10.10"
	member["stealth"] = true
	externalServers[0] = member

	computedExternalServers := BuildExternalServerListFromT(externalServers)
	assert.Equal(t, IBXExternalServersList, computedExternalServers)
}

func TestBuildExternalServerSetFromT(t *testing.T) {

	IBXExternalServersList := make([]common.ExternalServer, 1)
	b := true
	IBXExternalServer := common.ExternalServer{
		Address: "10.10.10.10",
		Name:    "Foo",
		Stealth: &b,
	}
	IBXExternalServersList[0] = IBXExternalServer

	member := make(map[string]interface{})
	member["name"] = "Foo"
	member["address"] = "10.10.10.10"
	member["stealth"] = true

	externalServers := &schema.Set{F: func(v interface{}) int {
		var buf bytes.Buffer
		value := v.(map[string]interface{})
		buf.Write([]byte(value["name"].(string)))
		buf.Write([]byte(value["address"].(string)))
		buf.Write([]byte(fmt.Sprintf("%t", value["stealth"].(bool))))
		return hashcode.String(buf.String())
	}}
	externalServers.Add(member)

	computedExternalServers := BuildExternalServerSetFromT(externalServers)
	assert.Equal(t, IBXExternalServersList, computedExternalServers)
}
