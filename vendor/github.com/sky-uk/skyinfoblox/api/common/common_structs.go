package common

// ExternalServer : external DNS server
type ExternalServer struct {
	Address                      string `json:"address"`
	Name                         string `json:"name"`
	SharedWithMSParentDelegation *bool  `json:"shared_with_ms_parent_delegation,omitempty"` //cannot be updated nor written
	Stealth                      *bool  `json:"stealth,omitempty"`                          //defaults to false
	TsigKey                      string `json:"tsig_key,omitempty"`                         //defaults to empty
	TsigKeyAlg                   string `json:"tsig_key_alg,omitempty"`                     // defaults to HMAC-MD5
	TsigKeyName                  string `json:"tsig_key_name,omitempty"`                    //defaults to empty
	UseTsigKeyName               *bool  `json:"use_tsig_key_name,omitempty"`                //defaults to false
}

// MemberServer : Grid member struct
type MemberServer struct {
	GridReplicate           *bool            `json:"grid_replicate,omitempty"`
	Lead                    *bool            `json:"lead,omitempty"`
	Name                    string           `json:"name,omitempty"`
	EnablePreferedPrimaries *bool            `json:"enable_preferred_primaries,omitempty"`
	PreferredPrimaries      []ExternalServer `json:"preferred_primaries,omitempty"`
	Stealth                 *bool            `json:"stealth,omitempty"`
}
