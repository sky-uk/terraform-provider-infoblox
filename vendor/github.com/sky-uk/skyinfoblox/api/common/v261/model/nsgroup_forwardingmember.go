package model

// NSGroupFwd : Name Server Group forwarding member object type
type NSGroupFwd struct {
	Reference         string                   `json:"_ref,omitempty"`
	Name              string                   `json:"name,omitempty"`
	Comment           string                   `json:"comment,omitempty"`
	ForwardingServers []ForwardingMemberServer `json:"forwarding_servers,omitempty"`
}
