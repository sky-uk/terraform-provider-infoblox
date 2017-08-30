package nsgroupfwd

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1"
const nsGroupFwdEndpoint = "/nsgroup:forwardingmember"

// RequestReturnFields : return fields used when making a request to the Infoblox API for this object type
var RequestReturnFields = []string{"name", "comment", "forwarding_servers"}

// NSGroupFwd : Name Server Group forwarding member object type
type NSGroupFwd struct {
	Reference         string                          `json:"_ref,omitempty"`
	Name              string                          `json:"name,omitempty"`
	Comment           string                          `json:"comment,omitempty"`
	ForwardingServers []common.ForwardingMemberServer `json:"forwarding_servers,omitempty"`
}
