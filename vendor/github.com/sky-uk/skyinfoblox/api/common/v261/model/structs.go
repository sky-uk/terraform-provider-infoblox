package model

// SchemaAttr - structures' attributes medatadata
type SchemaAttr struct {
	Type     string
	IsArray  bool
	Supports string
}

// StructAttrs - Get structure attribute metadata from schema
// Returns a map with attributes informations taken from schema
// this is manually compiled as there is no way to get struct schema
// from Infoblox
func StructAttrs() map[string]interface{} {
	s := map[string]interface{}{
		"extserver": map[string]SchemaAttr{
			"address": {Type: "string", IsArray: false, Supports: "rwu"},
			"name":    {Type: "string", IsArray: false, Supports: "rwu"},
			"shared_with_ms_parent_delegation": {Type: "string", IsArray: false, Supports: "r"},
			"stealth":                          {Type: "bool", IsArray: false, Supports: "rwu"},
			"tsig_key":                         {Type: "string", IsArray: false, Supports: "rwu"},
			"tsig_key_alg":                     {Type: "string", IsArray: false, Supports: "rwu"},
			"tsig_key_name":                    {Type: "string", IsArray: false, Supports: "rwu"},
		},
		"addressac": map[string]SchemaAttr{
			"_struct":    {Type: "string", IsArray: false, Supports: "rwu"},
			"address":    {Type: "string", IsArray: false, Supports: "rwu"},
			"permission": {Type: "string", IsArray: false, Supports: "rwu"},
		},
		"tsigac": map[string]SchemaAttr{
			"_struct":           {Type: "string", IsArray: false, Supports: "rwu"},
			"tsig_key":          {Type: "string", IsArray: false, Supports: "rwu"},
			"tsig_key_alg":      {Type: "string", IsArray: false, Supports: "rwu"},
			"tsig_key_name":     {Type: "string", IsArray: false, Supports: "rwu"},
			"use_tsig_key_name": {Type: "bool", IsArray: false, Supports: "rwu"},
		},
		"setting:scavenging": map[string]SchemaAttr{
			"ea_expression_list":          {Type: "array", IsArray: true, Supports: "rwu"},
			"enable_auto_reclamation":     {Type: "bool", IsArray: false, Supports: "rwu"},
			"enable_recurrent_scavenging": {Type: "bool", IsArray: false, Supports: "rwu"},
			"enable_rr_last_queried":      {Type: "bool", IsArray: false, Supports: "rwu"},
			"enable_scavenging":           {Type: "bool", IsArray: false, Supports: "rwu"},
			"enable_zone_last_queried":    {Type: "bool", IsArray: false, Supports: "rwu"},
			"expression_list":             {Type: "array", IsArray: false, Supports: "rwu"},
			"reclaim_associated_records":  {Type: "bool", IsArray: false, Supports: "rwu"},
			"scavenging_schedule":         {Type: "SchedulingSetting", IsArray: false, Supports: "rwu"},
		},
		"zonenameserver": map[string]SchemaAttr{
			"address":         {Type: "string", IsArray: false, Supports: "rwu"},
			"auto_create_ptr": {Type: "bool", IsArray: false, Supports: "rwu"},
		},
	}

	return s
}

// ExternalServer : external DNS server
type ExternalServer struct {
	Address                      string `json:"address"`
	Name                         string `json:"name"`
	SharedWithMSParentDelegation bool   `json:"shared_with_ms_parent_delegation,omitempty"` //cannot be updated nor written
	Stealth                      bool   `json:"stealth,omitempty"`                          //defaults to false
	TsigKey                      string `json:"tsig_key,omitempty"`                         //defaults to empty
	TsigKeyAlg                   string `json:"tsig_key_alg,omitempty"`                     // defaults to HMAC-MD5
	TsigKeyName                  string `json:"tsig_key_name,omitempty"`                    //defaults to empty
	UseTsigKeyName               bool   `json:"use_tsig_key_name,omitempty"`                //defaults to false
}

// MemberServer : Grid member struct
type MemberServer struct {
	GridReplicate           bool             `json:"grid_replicate,omitempty"`
	Lead                    bool             `json:"lead,omitempty"`
	Name                    string           `json:"name,omitempty"`
	EnablePreferedPrimaries bool             `json:"enable_preferred_primaries,omitempty"`
	PreferredPrimaries      []ExternalServer `json:"preferred_primaries,omitempty"`
	Stealth                 bool             `json:"stealth,omitempty"`
}

// ForwardingMemberServer - used by the zoneforward resource
type ForwardingMemberServer struct {
	Name                  string           `json:"name"`
	ForwardTo             []ExternalServer `json:"forward_to,omitempty"`
	ForwardersOnly        bool             `json:"forwarders_only,omitempty"`
	UseOverrideForwarders bool             `json:"use_override_forwarders,omitempty"`
}

// DHCPMember : Grid member serving DHCP
type DHCPMember struct {
	Ipv4addr string `json:"ipv4addr,omitempty"`
	Ipv6addr string `json:"ipv6addr,omitempty"`
	Name     string `json:"name,omitempty"`
}

// DHCPOption : set of options
type DHCPOption struct {
	Name        string `json:"name,omitempty"`
	Num         uint   `json:"num,omitempty"`
	UseOption   bool   `json:"use_option,omitempty"`
	Value       string `json:"value,omitempty"`
	VendorClass string `json:"vendor_class,omitempty"`
}

// ZoneAssociation : network association to a DNS zone
type ZoneAssociation struct {
	Fqdn      string `json:"fqdn"`
	IsDefault bool   `json:"is_default"`
	View      string `json:"view"`
}

// DNSSecKeyParameters : DNSSec Key Parameters
type DNSSecKeyParameters struct {
	EnableKskAutoRollover         bool                 `json:"enable_ksk_auto_rollover,omitempty"`
	KskAlgorithm                  string               `json:"ksk_algorithm,omitempty"`
	KskAlgorithms                 []DNSSecKeyAlgorithm `json:"ksk_algorithms,omitempty"`
	KskEmailNotificationEnabled   bool                 `json:"ksk_email_notification_enabled,omitempty"`
	KskRollover                   uint                 `json:"ksk_rollover,omitempty"`
	KskRolloverNotificationConfig string               `json:"ksk_rollover_notification_config,omitempty"`
	KskSize                       uint                 `json:"ksk_size,omitempty"`
	KskSnmpNotificationEnabled    bool                 `json:"ksk_snmp_notification_enabled,omitempty"`
	NextSecureType                string               `json:"next_secure_type,omitempty"`
	NSec3Iterations               uint                 `json:"nsec3_iterations,omitempty"`
	NSec3SaltMaxLength            uint                 `json:"nsec3_salt_max_length,omitempty"`
	NSec3SaltMinLength            uint                 `json:"nsec3_salt_min_length,omitempty"`
	SignatureExpiration           uint                 `json:"signature_expiration,omitempty"`
	ZskAlgorithmString            string               `json:"zsk_algorithm,omitempty"`
	ZskAlgorithms                 []ZskAlgorithm       `json:"zsk_algorithms,omitempty"`
	ZskRollover                   uint                 `json:"zsk_rollover,omitempty"`
	ZskRolloverMechanism          string               `json:"zsk_rollover_mechanism,omitempty"`
	ZskSize                       uint                 `json:"zsk_size,omitempty"`
}

// DNSSecKeyAlgorithm : algorithm structure for key signing and zone signing keys
type DNSSecKeyAlgorithm struct {
	Algorithm string `json:"algorithm,omitempty"`
	Size      uint   `json:"size,omitempty"`
}

// ZskAlgorithm : zone signing key algorithm
type ZskAlgorithm struct {
	Algorithm string `json:"algorithm,omitempty"`
	Size      uint   `json:"size,omitempty"`
}

// AddressAC : Access control rule for an address
type AddressAC struct {
	StructType string `json:"_struct,omitempty"`
	Address    string `json:"address,omitempty"`
	Permission string `json:"permission,omitempty"`
}

// TsigAC : TSIG key
type TsigAC struct {
	StructType     string `json:"_struct,omitemtpy"`
	TsigKey        string `json:"tsig_key,omitempty"`
	TsigKeyAlg     string `json:"tsig_key_alg,omitempty"`
	TsigKeyName    string `json:"tsig_key_name,omitempty"`
	UseTsigKeyName bool   `json:"use_tsig_key_name,omitempty"`
}

// AwsRte53ZoneInfo : Additional information for AWS Route53 zone
type AwsRte53ZoneInfo struct {
	AssociatedVPCs  []string `json:"associated_vpcs,omitempty,omitempty"`
	CallerReference string   `json:"caller_reference,omitempty"`
	DelegationSetID string   `json:"delegation_set_id,omitempty"`
	HostedZoneID    string   `json:"hosted_zone_id,omitempty"`
	NameServers     []string `json:"name_servers,omitempty"`
	RecordSetCount  uint     `json:"record_set_count,omitempty"`
	Type            string   `json:"type,omitempty"`
}

// CloudInformation : Contains Cloud API related information
type CloudInformation struct {
	DelegatedMember DHCPMember `json:"delegated_member,omitempty"`
	DelegatedRoot   string     `json:"delegated_root,omitempty"`
	DelegatedScope  string     `json:"delegated_scope,omitempty"`
	MGMTPlatform    string     `json:"mgmt_platform,omitempty"`
	OwnedByAdaptor  bool       `json:"owned_by_adaptor,omitempty"`
	Tenant          string     `json:"tenant,omitempty"`
	Usage           string     `json:"usage,omitempty"`
}

// DNSSecKey : DNS Sec Key
type DNSSecKey struct {
	Algorithm     string `json:"algorithm,omitempty"`
	NextEventDate string `json:"next_event_date,omitempty"`
	PublicKey     string `json:"public_key,omitempty"`
	Status        string `json:"status,omitempty"`
	Tag           uint   `json:"tag,omitempty"`
	Type          string `json:"type,omitempty"`
}

// SOAMName : SOA MNAME and primary server for the zone
type SOAMName struct {
	DNSMName        string `json:"dns_mname,omitempty"`
	GridPrimary     string `json:"grid_primary,omitempty"`
	MName           string `json:"mname,omitempty"`
	MSServerPrimary string `json:"ms_server_primary,omitempty"`
}

// GridMemberSOASerial : Grid Member per master SOA serial
type GridMemberSOASerial struct {
	GridPrimary     string `json:"grid_primary,omitempty"`
	MSServerPrimary string `json:"ms_server_primary,omitempty"`
	Serial          uint   `json:"serial,omitempty"`
}

// ADController : Active Directory controller object
type ADController struct {
	Address string `json:"address,omitempty"`
	Comment string `json:"comment,omitempty"`
}

// MSServer : Microsoft DNS server
type MSServer struct {
	Address                      string `json:"address,omitempty"`
	IsMaster                     bool   `json:"is_master,omitempty"`
	NSIP                         string `json:"ns_ip,omitempty"`
	NSName                       string `json:"ns_name,omitempty"`
	SharedWithMSParentDelegation bool   `json:"shared_with_ms_parent_delegation,omitempty"`
	Stealth                      bool   `json:"stealth,omitempty"`
}

// DNSZoneReference : A zone, it's reference and associated FQDN used for finding a zone when getting a list of all zones
type DNSZoneReference struct {
	Reference string `json:"_ref"`
	FQDN      string `json:"fqdn"`
}

// DNSScavengingSettings : Information about DNS Scavenging settings
type DNSScavengingSettings struct {
	EAExpressionList          []ExpressionOp  `json:"ea_expression_list,omitempty"`
	EnableAutoReclamation     bool            `json:"enable_auto_reclamation,omitempty"`
	EnableRecurrentScavenging bool            `json:"enable_recurrent_scavenging,omitempty"`
	EnableRRLastQueried       bool            `json:"enable_rr_last_queried,omitempty"`
	EnableScavenging          bool            `json:"enable_scavenging,omitempty"`
	EnableZoneLastQueried     bool            `json:"enable_zone_last_queried,omitempty"`
	ExpressionList            []ExpressionOp  `json:"expression_list,omitempty"`
	ReclaimAssociatedRecords  bool            `json:"reclaim_associated_records,omitempty"`
	ScavengingSchedule        ScheduleSetting `json:"scavenging_schedule,omitempty"`
}

// ScheduleSetting : Schedule settings
type ScheduleSetting struct {
	DayOfMonth      uint   `json:"day_of_month,omitempty"`
	Disable         bool   `json:"disable,omitempty"`
	Every           uint   `json:"every,omitempty"`
	Frequency       string `json:"frequency,omitempty"`
	HourOfDay       uint   `json:"hour_of_day,omitempty"`
	MinutesPastHour uint   `json:"minutes_past_hour,omitempty"`
	Month           uint   `json:"month,omitempty"`
	RecurringTime   string `json:"recurring_time,omitempty"`
	Repeat          string `json:"repeat,omitempty"`
	TimeZone        string `json:"time_zone,omitempty"`
	Weekdays        string `json:"weekdays,omitempty"`
	Year            uint   `json:"year,omitempty"`
}

// ExpressionOp : the extensible attribute operand structure
type ExpressionOp struct {
	OP      string `json:"op,omitempty"`
	OP1     string `json:"op1,omitempty"`
	OP1Type string `json:"op1_type,omitempty"`
	OP2     string `json:"op2,omitempty"`
	OP2Type string `json:"op2_type,omitempty"`
}

//ZoneNameServer : The Zone Name Server structure provides IP address information for the name server associated with a NS record
type ZoneNameServer struct {
	Address                 string `json:"address,omitempty"`
	AutoCreatePointerRecord bool   `json:"auto_create_ptr,omitempty"`
}

// DNSSecTrustedKey : This is the the DNSKEY record that holds the KSK as a trust anchor for each zone for which the Grid member returns validated data.
type DNSSecTrustedKey struct {
	Algorithim       string `json:"algorithm,omitempty"`
	FQDN             string `json:"fqdn,omitempty"`
	Key              string `json:"key,omitempty"`
	SecureEntryPoint bool   `json:"secure_entry_point,omitempty"`
}

// FixedRRSetOrderFQDN : A fixed RRset order FQDN contains information about the fixed RRset configuration items.
type FixedRRSetOrderFQDN struct {
	FQDN       string `json:"fqdn,omitempty"`
	RecordType string `json:"record_type,omitempty"`
}

// DNSResponseRateLimiting : The DNS Response Rate Limiting structure provides information about DNS response rate limiting configuration.
type DNSResponseRateLimiting struct {
	EnableRRL          bool `json:"enable_rrl,omitempty"`
	LogOnly            bool `json:"log_only,omitempty"`
	ResponsesPerSecond uint `json:"responses_per_second,omitempty"`
	Slip               uint `json:"slip,omitempty"`
	Window             uint `json:"window,omitempty"`
}

// DNSSortlist : A sortlist defines the order of IP addresses listed in responses sent to DNS queries.
type DNSSortlist struct {
	Address   string   `json:"address,omitempty"`
	MatchList []string `json:"match_list,omitempty"`
}
