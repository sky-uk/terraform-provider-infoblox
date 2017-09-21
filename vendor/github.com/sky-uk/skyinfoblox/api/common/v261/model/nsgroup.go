package model

// NSGroupAuth : Name Server Group Authoritative object type
type NSGroupAuth struct {
	Reference           string           `json:"_ref,omitempty"`
	Comment             string           `json:"comment,omitempty"`
	ExternalPrimaries   []ExternalServer `json:"external_primaries,omitempty"`
	ExternalSecondaries []ExternalServer `json:"external_secondaries,omitempty"`
	GridPrimary         []MemberServer   `json:"grid_primary,omitempty"`
	GridSecondaries     []MemberServer   `json:"grid_secondaries,omitempty"`
	GridDefault         bool             `json:"is_grid_default,omitempty"`
	Name                string           `json:"name,omitempty"`
	UseExternalPrimary  bool             `json:"use_external_primary,omitempty"`
}
