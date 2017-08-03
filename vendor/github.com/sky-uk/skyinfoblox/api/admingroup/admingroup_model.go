package admingroup

const adminGroupEndpoint = "/wapi/v2.3.1"

// IBXAdminGroup : Admin group definition
type IBXAdminGroup struct {
	Reference      string   `json:"_ref,omitempty"`
	AccessMethod   []string `json:"access_method,omitempty"`
	Comment        string   `json:"comment,omitempty"`
	Disable        *bool    `json:"disable,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	Name           string   `json:"name,omitempty"`
	Roles          []string `json:"roles,omitempty"`
	SuperUser      *bool    `json:"superuser,omitempty"`
}

// IBXAdminGroupReference : A reference object for an admin group
type IBXAdminGroupReference struct {
	Reference      string `json:"_ref"`
	AdminGroupName string `json:"name"`
}
