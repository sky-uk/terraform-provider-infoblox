package adminuser

// AdminUser struct
type AdminUser struct {
	Ref      string   `json:"_ref"`
	Name     string   `json:"name"`
	Groups   []string `json:"admin_groups"`
	Email    string   `json:"email,omitempty"`
	Disable  *bool    `json:"disable,omitempty"`
	Comment  string   `json:"comment,omitempty"`
	Password string   `json:"password"`
}
