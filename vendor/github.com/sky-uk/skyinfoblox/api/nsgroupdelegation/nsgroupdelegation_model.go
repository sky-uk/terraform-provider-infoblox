package nsgroupdelegation

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1"
const nsGroupDelegationEndpoint = "/nsgroup:delegation"

// RequestReturnFields : return fields used when making a request to the Infoblox API for this object type
var RequestReturnFields = []string{"comment", "name", "delegate_to"}

// NSGroupDelegation : Name Server Group Delegation object type
type NSGroupDelegation struct {
	Reference  string                  `json:"_ref,omitempty"`
	Comment    string                  `json:"comment,omitempty"`
	DelegateTo []common.ExternalServer `json:"delegate_to,omitempty"`
	Name       string                  `json:"name,omitempty"`
}
