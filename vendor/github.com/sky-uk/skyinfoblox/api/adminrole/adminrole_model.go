package adminrole

// AdminRole struct
type AdminRole struct {
	Reference string `json:"_ref,omitempty"`
	Name      string `json:"name,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Disable   *bool  `json:"disable,omitempty"`
}
