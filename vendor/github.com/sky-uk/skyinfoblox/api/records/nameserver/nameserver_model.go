package nameserver

const wapiVersion = "/wapi/v2.6.1"
const nsEndpoint = "/record:ns"

// NSRecord : Name server record
type NSRecord struct {
	Reference        string           `json:"_ref,omitempty"`
	Name             string           `json:"name,omitempty"`
	Addresses        []ZoneNameServer `json:"addresses,omitempty"`
	MSDelegationName string           `json:"ms_delegation_name,omitempty"`
	NameServer       string           `json:"nameserver,omitempty"`
	View             string           `json:"view,omitempty"`
}

//ZoneNameServer : The Zone Name Server structure provides IP address information for the name server associated with a NS record
type ZoneNameServer struct {
	Address                 string `json:"address,omitempty"`
	AutoCreatePointerRecord *bool  `json:"auto_create_ptr,omitempty"`
}
