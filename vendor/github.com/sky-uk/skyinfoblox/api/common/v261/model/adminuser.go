package model

// AdminUser struct
type AdminUser struct {
	AdminGroups []string `json:"admin_groups"`
	AuthType    string   `json:"auth_type,omitempty"` //enum: valid values("LOCAL","REMOTE"), defaults to "LOCAL"
	Comment     string   `json:"comment,omitempty"`
	Disable     bool     `json:"disable,omitempty"`
	Email       string   `json:"email,omitempty"`
	Name        string   `json:"name"`
	Password    string   `json:"password"`
	TimeZone    string   `json:"time_zone,omitempty"`     //enum: valid values("(UTC) [+/- <time>]"), defaults to "(UTC)"
	UseTimeZone bool     `json:"use_time_zone,omitempty"` // defaults to false
}
