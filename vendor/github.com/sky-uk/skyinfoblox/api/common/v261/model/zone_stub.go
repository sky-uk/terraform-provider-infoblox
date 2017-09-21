package model

// ZoneStub - default struct for the stub zone
type ZoneStub struct {
	Ref                string           `json:"_ref,omitempty"`
	Comment            string           `json:"comment,omitempty"`
	Disable            bool             `json:"disable,omitempty"`
	DisableForwarding  bool             `json:"disable_forwarding,omitempty"`
	ExternalNSGroup    string           `json:"external_ns_group,omitempty"`
	FQDN               string           `json:"fqdn,omitempty"`
	Locked             bool             `json:"locked,omitempty"`
	MaskPrefix         string           `json:"mask_prefix,omitempty"`
	NsGroup            string           `json:"ns_group,omitempty"`
	Prefix             string           `json:"prefix,omitempty"`
	StubFrom           []ExternalServer `json:"stub_from,omitempty"`
	StubMembers        []MemberServer   `json:"stub_members,omitempty"`
	UseSRGAssociations bool             `json:"using_srg_associations,omitempty"`
	View               string           `json:"view,omitempty"`
	ZoneFormat         string           `json:"zone_format,omitempty"`
}
