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

// ForwardingMemberServer - used by the zoneforward resource
type ForwardingMemberServer struct {
	Name                  string           `json:"name"`
	ForwardTo             []ExternalServer `json:"forward_to,omitempty"`
	ForwardersOnly        *bool            `json:"forwarders_only,omitempty"`
	UseOverrideForwarders *bool            `json:"use_override_forwarders,omitempty"`
}

// DNSSecTrustedKey : This is the the DNSKEY record that holds the KSK as a trust anchor for each zone for which the Grid member returns validated data.
type DNSSecTrustedKey struct {
	Algorithim       string `json:"algorithm,omitempty"`
	FQDN             string `json:"fqdn,omitempty"`
	Key              string `json:"key,omitempty"`
	SecureEntryPoint *bool  `json:"secure_entry_point,omitempty"`
}

// AddressAC : This struct represents an access control rule for an address.
type AddressAC struct {
	Address    string `json:"address,omitempty"`
	Permission string `json:"permission,omitempty"`
}

// TsigAC : This struct represents a TSIG key.
type TsigAC struct {
	TsigKey        string `json:"tsig_key,omitempty"`
	TsigKeyAlg     string `json:"tsig_key_alg,omitempty"`
	TsigKeyName    string `json:"tsig_key_name,omitempty"`
	UseTsigKeyName *bool  `json:"use_tsig_key_name,omitempty"`
}

// FixedRRSetOrderFQDN : A fixed RRset order FQDN contains information about the fixed RRset configuration items.
type FixedRRSetOrderFQDN struct {
	FQDN       string `json:"fqdn,omitempty"`
	RecordType string `json:"record_type,omitempty"`
}

// DNSResponseRateLimiting : The DNS Response Rate Limiting structure provides information about DNS response rate limiting configuration.
type DNSResponseRateLimiting struct {
	EnableRRL          *bool `json:"enable_rrl,omitempty"`
	LogOnly            *bool `json:"log_only,omitempty"`
	ResponsesPerSecond uint  `json:"responses_per_second,omitempty"`
	Slip               uint  `json:"slip,omitempty"`
	Window             uint  `json:"window,omitempty"`
}

// DNSScavengingSettings : The DNS scavenging settings object provides information about scavenging configuration e.g. conditions under which records can be scavenged, periodicity of scavenging operations.
type DNSScavengingSettings struct {
	EaExpressionList          []ExpressionOperand `json:"ea_expression_list,omitempty"`
	EnableAutoReclamation     *bool               `json:"enable_auto_reclamation,omitempty"`
	EnableRecurrentScavenging *bool               `json:"enable_recurrent_scavenging,omitempty"`
	EnableRRLastQueried       *bool               `json:"enable_rr_last_queried,omitempty"`
	EnableScavenging          *bool               `json:"enable_scavenging,omitempty"`
	EnableZoneLastQueried     *bool               `json:"enable_zone_last_queried,omitempty"`
	ExpressionList            []ExpressionOperand `json:"expression_list,omitempty"`
	ReclaimAssociatedRecords  *bool               `json:"reclaim_associated_records,omitempty"`
	ScavengingSchedule        ScheduleSetting     `json:"scavenging_schedule,omitempty"`
}

// ExpressionOperand : The expression operand structure is used to build expression lists. The allowed values for the expression operand structure depend on the object they appear to be a part of.
type ExpressionOperand struct {
	Op      string `json:"op,omitempty"`
	Op1     string `json:"op1,omitempty"`
	Op1Type string `json:"op1_type,omitempty"`
	Op2     string `json:"op2,omitempty"`
	Op2Type string `json:"op2_type,omitempty"`
}

// ScheduleSetting : This struct contains information about scheduling settings.
type ScheduleSetting struct {
	DayOfMonth      uint     `json:"day_of_month,omitempty"`
	Disable         *bool    `json:"disable,omitempty"`
	Every           uint     `json:"every,omitempty"`
	Frequency       string   `json:"frequency,omitempty"`
	HourOfDay       uint     `json:"hour_of_day,omitempty"`
	MinutesPastHour uint     `json:"minutes_past_hour,omitempty"`
	Month           uint     `json:"month,omitempty"`
	Repeat          string   `json:"repeat,omitempty"`
	TimeZone        string   `json:"time_zone,omitempty"`
	Weekdays        []string `json:"weekdays,omitempty"`
	Year            uint     `json:"year,omitempty"`
}

// DNSSortlist : A sortlist defines the order of IP addresses listed in responses sent to DNS queries.
type DNSSortlist struct {
	Address   string   `json:"address,omitempty"`
	MatchList []string `json:"match_list,omitempty"`
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
	UseOption   *bool  `json:"use_option,omitempty"`
	Value       string `json:"value,omitempty"`
	VendorClass string `json:"vendor_class,omitempty"`
}

// ZoneAssociation : network association to a DNS zone
type ZoneAssociation struct {
	Fqdn      string `json:"fqdn"`
	IsDefault bool   `json:"is_default"`
	View      string `json:"view"`
}
