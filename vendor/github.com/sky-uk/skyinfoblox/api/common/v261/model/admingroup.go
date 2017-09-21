package model

// AdminGroup : Admin group definition
type AdminGroup struct {
	AccessMethod   []string `json:"access_method,omitempty"`
	Comment        string   `json:"comment,omitempty"`
	Disable        bool     `json:"disable,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	Name           string   `json:"name"`
	Roles          []string `json:"roles,omitempty"`
	SuperUser      bool     `json:"superuser,omitempty"`
}
