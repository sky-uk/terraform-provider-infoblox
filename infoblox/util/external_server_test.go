package util

import (
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
