package model

// AdminRole struct
type AdminRole struct {
	Name    string `json:"name"`
	Comment string `json:"comment,omitempty"`
	Disable bool   `json:"disable,omitempty"`
}
