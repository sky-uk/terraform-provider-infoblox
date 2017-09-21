package model

// NSGroupFwdStub : Name Server Group forward/stub object type
type NSGroupFwdStub struct {
	Reference       string           `json:"_ref,omitempty"`
	Name            string           `json:"name,omitempty"`
	Comment         string           `json:"comment,omitempty"`
	ExternalServers []ExternalServer `json:"external_servers,omitempty"`
}
