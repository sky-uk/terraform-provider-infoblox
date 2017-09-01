package nsgroupfwdstub

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1"
const nsGroupFwdStubEndpoint = "/nsgroup:forwardstubserver"

// RequestReturnFields : return fields used when making a request to the Infoblox API for this object type
var RequestReturnFields = []string{"name", "comment", "external_servers"}

// NSGroupFwdStub : Name Server Group forward/stub object type
type NSGroupFwdStub struct {
	Reference       string                  `json:"_ref,omitempty"`
	Name            string                  `json:"name,omitempty"`
	Comment         string                  `json:"comment,omitempty"`
	ExternalServers []common.ExternalServer `json:"external_servers,omitempty"`
}
