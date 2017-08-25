package dhcprange

const wapiVersion = "/wapi/v2.6.1"

// DHCPRange struct
type DHCPRange struct {
	Ref               string `json:"_ref"`
	Start             string `json:"start_addr"`
	End               string `json:"end_addr"`
	Network           string `json:"network"`
	NetworkView       string `json:"network_view"`
	Restart           *bool  `json:"restart_if_needed,omitempty"`
	ServerAssociation string `json:"server_association_type,omitempty"`
	Name              string `json:"name,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Member            Member `json:"member,omitempty"`
}

// Member - Grid member serving DHCP struct
// All members in the array must be of the same type.
// the struct type must be indicated in each element, by setting the “_struct” member to the struct type.
type Member struct {
	ElementType string `json:"_struct"`
	IPv4Address string `json:"ipv4addr,omitempty"`
	IPv6Address string `json:"ipv6addr,omitempty"`
	Name        string `json:"name,omitempty"`
}
