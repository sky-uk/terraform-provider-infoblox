package nsgroupstub

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1"
const nsGroupStubEndpoint = "/nsgroup:stubmember"

// RequestReturnFields : return fields used when making a request to the Infoblox API for this object type
var RequestReturnFields = []string{"name", "comment", "stub_members"}

// NSGroupStub : Name Server Group stub object type
type NSGroupStub struct {
	Reference   string                `json:"_ref,omitempty"`
	Name        string                `json:"name,omitempty"`
	Comment     string                `json:"comment,omitempty"`
	StubMembers []common.MemberServer `json:"stub_members,omitempty"`
}
