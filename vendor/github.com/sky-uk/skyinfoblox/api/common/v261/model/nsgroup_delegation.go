package model

// NSGroupDelegation : Name Server Group Delegation object type
type NSGroupDelegation struct {
	Comment    string           `json:"comment,omitempty"`
	DelegateTo []ExternalServer `json:"delegate_to,omitempty"`
	Name       string           `json:"name,omitempty"`
}
