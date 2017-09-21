package model

// ZoneDelegated - Main struct for zone delegation
type ZoneDelegated struct {
	Ref                    string           `json:"_ref,omitempty"`
	Address                string           `json:"address,omitempty"`
	Comment                string           `json:"comment,omitempty"`
	DelegateTo             []ExternalServer `json:"delegate_to,omitempty"`
	DelegatedTTL           uint             `json:"delegated_ttl,omitempty"`
	Disable                *bool            `json:"disable,omitempty"`
	DNSFqdn                string           `json:"dns_fqdn,omitempty"`
	EnableRFC2317Exclusion *bool            `json:"enable_rfc2317_exclusion,omitempty"`
	Fqdn                   string           `json:"fqdn,omitempty"`
	Locked                 *bool            `json:"locked,omitempty"`
	Prefix                 string           `json:"prefix,omitempty"`
	UseDelegatedTTL        *bool            `json:"use_delegated_ttl,omitempty"`
	View                   string           `json:"view,omitempty"`
	ZoneFormat             string           `json:"zone_format,omitempty"`
	NsGroup                string           `json:"ns_group,omitempty"`
}
