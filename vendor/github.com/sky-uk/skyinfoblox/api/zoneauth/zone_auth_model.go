package zoneauth

import (
	"github.com/sky-uk/skyinfoblox/api/common"
)

/* Notes:
AllowQuery, AllowTransfer & AllowUpdate can be either AddressAC or TsigAC
Can't send this empty as it causes an error 'AwsRte53ZoneInfoList                    AwsRte53ZoneInfo      `json:"aws_rte53_zone_info,omitempty"`'
NetworkAssociations is an array of network, networkcontainer, ipv6network, ipv6networkcontainer - can't be written or updated.
UpdateForwarding can be one of the following: Address ac struct, TSIG ac struct array. */

// DNSZone : Contains zone configuration. Reference is used during updates and when retriving the zone.
type DNSZone struct {
	Reference                               string                  `json:"_ref,omitempty"`
	FQDN                                    string                  `json:"fqdn,omitempty"`
	View                                    string                  `json:"view,omitempty"`
	Comment                                 string                  `json:"comment,omitempty"`
	Address                                 string                  `json:"address,omitempty"`
	AllowActiveDir                          []interface{}           `json:"allow_active_dir,omitempty"`
	AllowGssTsigUnderScoreZone              *bool                   `json:"allow_gss_tsig_for_underscore_zone,omitempty"`
	AllowGssTsigZoneUpdates                 *bool                   `json:"allow_gss_tsig_zone_updates,omitempty"`
	AllowQuery                              []interface{}           `json:"allow_query,omitempty"`
	AllowTransfer                           []interface{}           `json:"allow_transfer,omitempty"`
	AllowUpdate                             []interface{}           `json:"allow_update,omitempty"`
	AllowUpdateForwarding                   *bool                   `json:"allow_update_forwarding,omitempty"`
	CloudInfo                               []CloudInformation      `json:"cloud_info,omitempty"`
	CopyXferToNotify                        *bool                   `json:"copy_xfer_to_notify,omitempty"`
	CreatePtrBulkHosts                      *bool                   `json:"create_ptr_for_bulk_hosts,omitempty"`
	CreatePtrHosts                          *bool                   `json:"create_ptr_for_hosts,omitempty"`
	CreateUnderscoreZones                   *bool                   `json:"create_underscore_zones,omitempty"`
	DDNSPrincipleGroup                      string                  `json:"ddns_principal_group,omitempty"`
	DDNSPrincipleTracking                   *bool                   `json:"ddns_principal_tracking,omitempty"`
	DDNSRestrictPatterns                    *bool                   `json:"ddns_restrict_patterns,omitempty"`
	DDNSRestrictPatternsList                []string                `json:"ddns_restrict_patterns_list,omitempty"`
	DDNSRestrictProtected                   *bool                   `json:"ddns_restrict_protected,omitempty"`
	DDNSRestrictSecure                      *bool                   `json:"ddns_restrict_secure,omitempty"`
	DDNSRestrictStatic                      *bool                   `json:"ddns_restrict_static,omitempty"`
	Disable                                 *bool                   `json:"disable,omitempty"`
	DisableForwarding                       *bool                   `json:"disable_forwarding,omitempty"`
	DisplayDomain                           string                  `json:"display_domain,omitempty"`
	DNSFqdn                                 string                  `json:"dns_fqdn,omitempty"`
	DNSIntegrityEnable                      *bool                   `json:"dns_integrity_enable,omitempty"`
	DNSIntegrityFrequency                   uint                    `json:"dns_integrity_frequency,omitempty"`
	DNSIntegrityMember                      string                  `json:"dns_integrity_member,omitempty"`
	DNSIntegrityVerboseLogging              *bool                   `json:"dns_integrity_verbose_logging,omitempty"`
	DNSSoaEmail                             string                  `json:"dns_soa_email,omitempty"`
	DNSSecKeyParams                         []DNSSecKeyParameters   `json:"dnssec_key_params,omitempty"`
	DNSSecKeys                              []DNSSecKey             `json:"dnssec_keys,omitempty"`
	DNSSecKskRolloverDate                   string                  `json:"dnssec_ksk_rollover_date,omitempty"`
	DNSSecZskRolloverDate                   string                  `json:"dnssec_zsk_rollover_date,omitempty"`
	DoHostAbstraction                       *bool                   `json:"do_host_abstraction,omitempty"`
	EffectiveCheckNamesPolicy               string                  `json:"effective_check_names_policy,omitempty"`
	EffectiveRecordNamePolicy               string                  `json:"effective_record_name_policy,omitempty"`
	ExtAttrs                                string                  `json:"extattrs,omitempty"`
	ExternalPrimaries                       []common.ExternalServer `json:"external_primaries,omitempty"`
	ExternalSecondaries                     []common.ExternalServer `json:"external_secondaries,omitempty"`
	GridPrimary                             []common.MemberServer   `json:"grid_primary,omitempty"`
	GridPrimarySharedWithMSParentDelegation *bool                   `json:"grid_primary_shared_with_ms_parent_delegation,omitempty"`
	GridSecondaries                         []common.MemberServer   `json:"grid_secondaries,omitempty"`
	ImportFrom                              string                  `json:"import_from,omitempty"`
	IsDNSSecEnabled                         *bool                   `json:"is_dnssec_enabled,omitempty"`
	IsDNSSecSigned                          *bool                   `json:"is_dnssec_signed,omitempty"`
	IsMultiMaster                           *bool                   `json:"is_multimaster,omitempty"`
	LastQueried                             string                  `json:"last_queried,omitempty"`
	Locked                                  *bool                   `json:"locked,omitempty"`
	LockedBy                                string                  `json:"locked_by,omitempty"`
	MaskPrefix                              string                  `json:"mask_prefix,omitempty"`
	MemberSOAMNames                         []SOAMName              `json:"member_soa_mnames,omitempty"`
	MemberSOASerials                        []GridMemberSOASerial   `json:"member_soa_serials,omitempty"`
	MSADIntegrated                          *bool                   `json:"ms_ad_integrated,omitempty"`
	MSAllowTransfer                         []interface{}           `json:"ms_allow_transfer,omitempty"`
	MSAllowTransferMode                     string                  `json:"ms_allow_transfer_mode,omitempty"`
	MSDCNSRecordCreation                    []ADController          `json:"ms_dc_ns_record_creation,omitempty"`
	MSDDNSMode                              string                  `json:"ms_ddns_mode,omitempty"`
	MSManaged                               string                  `json:"ms_managed,omitempty"`
	MSPrimaries                             []MSServer              `json:"ms_primaries,omitempty"`
	MSReadOnly                              *bool                   `json:"ms_read_only,omitempty"`
	MSSecondaries                           []MSServer              `json:"ms_secondaries,omitempty"`
	MSSyncDisabled                          *bool                   `json:"ms_sync_disabled,omitempty"`
	MSSyncMasterName                        string                  `json:"ms_sync_master_name,omitempty"`
	NetworkAssociations                     []string                `json:"network_associations,omitempty"`
	NetworkView                             string                  `json:"network_view,omitempty"`
	NotifyDelay                             uint                    `json:"notify_delay,omitempty"`
	NSGroup                                 string                  `json:"ns_group,omitempty"`
	Parent                                  string                  `json:"parent,omitempty"`
	Prefix                                  string                  `json:"prefix,omitempty"`
	PrimaryType                             string                  `json:"primary_type,omitempty"`
	RecordNamePolicy                        string                  `json:"record_name_policy,omitempty"`
	RecordsMonitored                        *bool                   `json:"records_monitored,omitempty"`
	RestartIfNeeded                         *bool                   `json:"restart_if_needed,omitempty"`
	RRNotQueriedEnabledTime                 string                  `json:"rr_not_queried_enabled_time,omitempty"`
	ScavengingSettings                      DNSScavengingSettings   `json:"scavenging_settings,omitempty"`
	SetSOASerialNumber                      *bool                   `json:"set_soa_serial_number,omitempty"`
	SOADefaultTTL                           uint                    `json:"soa_default_ttl,omitempty"`
	SOAEmail                                string                  `json:"soa_email,omitempty"`
	SOAExpire                               uint                    `json:"soa_expire,omitempty"`
	SOANegativeTTL                          uint                    `json:"soa_negative_ttl,omitempty"`
	SOARefresh                              uint                    `json:"soa_refresh,omitempty"`
	SOARetry                                uint                    `json:"soa_retry,omitempty"`
	SOASerialNumber                         uint                    `json:"soa_serial_number,omitempty"`
	SRGS                                    string                  `json:"srgs,omitempty"`
	UpdateForwarding                        []interface{}           `json:"update_forwarding,omitempty"`
	UseAllowActiveDir                       *bool                   `json:"use_allow_active_dir,omitempty"`
	UseAllowQuery                           *bool                   `json:"use_allow_query,omitempty"`
	UseAllowTransfer                        *bool                   `json:"use_allow_transfer,omitempty"`
	UseAllowUpdate                          *bool                   `json:"use_allow_update,omitempty"`
	UseAllowUpdateForwarding                *bool                   `json:"use_allow_update_forwarding,omitempty"`
	UseCheckNamesPolicy                     *bool                   `json:"use_check_names_policy,omitempty"`
	UseCopyXferNotify                       *bool                   `json:"use_copy_xfer_to_notify,omitempty"`
	UseDDNSPatternsRestriction              *bool                   `json:"use_ddns_patterns_restriction,omitempty"`
	UseDDNSPrincipleSecurity                *bool                   `json:"use_ddns_principal_security,omitempty"`
	UseDDNSRestrictProtected                *bool                   `json:"use_ddns_restrict_protected,omitempty"`
	UseDDNSrestrictStatic                   *bool                   `json:"use_ddns_restrict_static,omitempty"`
	UseDDNSSecKeyParams                     *bool                   `json:"use_dnssec_key_params,omitempty"`
	UseExternalPrimary                      *bool                   `json:"use_external_primary,omitempty"`
	UseGridZoneTimer                        *bool                   `json:"use_grid_zone_timer,omitempty"`
	UseImportFrom                           *bool                   `json:"use_import_from,omitempty"`
	UseNotifyDelay                          *bool                   `json:"use_notify_delay,omitempty"`
	UseRecordNamePolicy                     *bool                   `json:"use_record_name_policy,omitempty"`
	UseScavengingSettings                   *bool                   `json:"use_scavenging_settings,omitempty"`
	UseSOAEmail                             *bool                   `json:"use_soa_email,omitempty"`
	UsingSrgAssociations                    *bool                   `json:"using_srg_associations,omitempty"`
	ZoneFormat                              string                  `json:"zone_format,omitempty"`
	ZoneNotQueriedEnabledTime               string                  `json:"zone_not_queried_enabled_time,omitempty"`
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
	UseTsigKeyName *bool  `json:"use_tsig_key_name,omitempty"`
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
	OwnedByAdaptor  *bool      `json:"owned_by_adaptor,omitempty"`
	Tenant          string     `json:"tenant,omitempty"`
	Usage           string     `json:"usage,omitempty"`
}

// DHCPMember : contains names and addresses of DHCP serving grid members
type DHCPMember struct {
	IPV4Addr string `json:"ipv4addr,omitempty"`
	IPV6Addr string `json:"ipv6addr,omitempty"`
	Name     string `json:"name,omitempty"`
}

// DNSSecKeyParameters : DNSSec Key Parameters
type DNSSecKeyParameters struct {
	EnableKskAutoRollover         *bool                `json:"enable_ksk_auto_rollover,omitempty"`
	KskAlgorithm                  string               `json:"ksk_algorithm,omitempty"`
	KskAlgorithms                 []DNSSecKeyAlgorithm `json:"ksk_algorithms,omitempty"`
	KskEmailNotificationEnabled   *bool                `json:"ksk_email_notification_enabled,omitempty"`
	KskRollover                   uint                 `json:"ksk_rollover,omitempty"`
	KskRolloverNotificationConfig string               `json:"ksk_rollover_notification_config,omitempty"`
	KskSize                       uint                 `json:"ksk_size,omitempty"`
	KskSnmpNotificationEnabled    *bool                `json:"ksk_snmp_notification_enabled,omitempty"`
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

// ZskAlgorithm : zone signing key algorithm
type ZskAlgorithm struct {
	Algorithm string `json:"algorithm,omitempty"`
	Size      uint   `json:"size,omitempty"`
}

// DNSSecKeyAlgorithm : algorithm structure for key signing and zone signing keys
type DNSSecKeyAlgorithm struct {
	Algorithm string `json:"algorithm,omitempty"`
	Size      uint   `json:"size,omitempty"`
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
	IsMaster                     *bool  `json:"is_master,omitempty"`
	NSIP                         string `json:"ns_ip,omitempty"`
	NSName                       string `json:"ns_name,omitempty"`
	SharedWithMSParentDelegation *bool  `json:"shared_with_ms_parent_delegation,omitempty"`
	Stealth                      *bool  `json:"stealth,omitempty"`
}

// DNSZoneReference : A zone, it's reference and associated FQDN used for finding a zone when getting a list of all zones
type DNSZoneReference struct {
	Reference string `json:"_ref"`
	FQDN      string `json:"fqdn"`
}

// DNSScavengingSettings : Information about DNS Scavenging settings
type DNSScavengingSettings struct {
	EAExpressionList          []ExpressionOp  `json:"ea_expression_list,omitempty"`
	EnableAutoReclamation     *bool           `json:"enable_auto_reclamation,omitempty"`
	EnableRecurrentScavenging *bool           `json:"enable_recurrent_scavenging,omitempty"`
	EnableRRLastQueried       *bool           `json:"enable_rr_last_queried,omitempty"`
	EnableScavenging          *bool           `json:"enable_scavenging,omitempty"`
	EnableZoneLastQueried     *bool           `json:"enable_zone_last_queried,omitempty"`
	ExpressionList            []ExpressionOp  `json:"expression_list,omitempty"`
	ReclaimAssociatedRecords  *bool           `json:"reclaim_associated_records,omitempty"`
	ScavengingSchedule        ScheduleSetting `json:"scavenging_schedule,omitempty"`
}

// ScheduleSetting : Schedule settings
type ScheduleSetting struct {
	DayOfMonth      uint   `json:"day_of_month,omitempty"`
	Disable         *bool  `json:"disable,omitempty"`
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

// DNSZoneReferences : A list of zone references
type DNSZoneReferences []DNSZoneReference
