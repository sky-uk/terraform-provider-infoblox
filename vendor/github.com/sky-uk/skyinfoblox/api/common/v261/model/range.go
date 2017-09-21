package model

// Range - DHCP Range struct
type Range struct {
	Ref               string     `json:"_ref"`
	Start             string     `json:"start_addr"`
	End               string     `json:"end_addr"`
	Network           string     `json:"network"`
	NetworkView       string     `json:"network_view"`
	Restart           bool       `json:"restart_if_needed,omitempty"`
	ServerAssociation string     `json:"server_association_type,omitempty"`
	Name              string     `json:"name,omitempty"`
	Comment           string     `json:"comment,omitempty"`
	Member            DHCPMember `json:"member,omitempty"`
}
