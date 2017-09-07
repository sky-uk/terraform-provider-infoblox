package mxrecord

//WapiVersion : WAPI version related with this data model
const wapiVersion = "/wapi/v2.6.1"

// MXRecordEndpoint - resource WAPI endpoint
const mxRecordEndpoint = "record:mx"

// MxRecord struct
type MxRecord struct {
	Ref               string `json:"_ref,omitempty"`
	Comment           string `json:"comment,omitempty"`
	DDNSPrincipal     string `json:"ddns_principal,omitempty"`
	DDNSProtected     bool   `json:"ddns_protected"`
	Disable           bool   `json:"disable"`
	ForbidReclamation bool   `json:"forbid_reclamation"`
	MailExchanger     string `json:"mail_exchanger"`
	Name              string `json:"name"`
	Preference        uint   `json:"preference"`
	TTL               uint   `json:"ttl,omitempty"`
	UseTTL            bool   `json:"use_ttl"`
	View              string `json:"view,omitempty"`
}
