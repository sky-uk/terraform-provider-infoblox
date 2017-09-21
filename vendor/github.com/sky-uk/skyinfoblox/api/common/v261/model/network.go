package model

// Network : base DHCP Network object model
type Network struct {
	Ref                              string            `json:"_ref"`
	Network                          string            `json:"network,omitempty"`
	NetworkView                      string            `json:"network_view,omitempty"`
	Comment                          string            `json:"comment,omitempty"`
	Authority                        bool              `json:"authority,omitempty"`
	AutoCreateReversezone            bool              `json:"auto_create_reversezone,omitempty"`
	Disable                          bool              `json:"disable,omitempty"`
	EnableDdns                       bool              `json:"enable_ddns,omitempty"`
	EnableDhcpThresholds             bool              `json:"enable_dhcp_thresholds,omitempty"`
	HighWaterMark                    int               `json:"high_water_mark,omitempty"`
	HighWaterMarkReset               int               `json:"high_water_mark_reset,omitempty"`
	LowWaterMark                     int               `json:"low_water_mark,omitempty"`
	LowWaterMarkReset                int               `json:"low_water_mark_reset,omitempty"`
	EnableDiscovery                  bool              `json:"enable_discovery,omitempty"`
	DiscoveryMember                  string            `json:"discovery_member,omitempty"`
	Ipv4addr                         string            `json:"ipv4addr,omitempty"`
	LeaseScavengeTime                int               `json:"lease_scavenge_time,omitempty"`
	Netmask                          uint              `json:"netmask,omitempty"`
	NetworkContainer                 string            `json:"network_container,omitempty"`
	Options                          []DHCPOption      `json:"options,omitempty"`
	Members                          []DHCPMember      `json:"members,omitempty"`
	RecycleLeases                    bool              `json:"recycle_leases,omitempty"`
	RestartIfNeeded                  bool              `json:"restart_if_needed,omitempty"`
	UpdateDNSOnLeaseRenewal          bool              `json:"update_dns_on_lease_renewal,omitempty"`
	UseAuthority                     bool              `json:"use_authority,omitempty"`
	UseBlackoutSetting               bool              `json:"use_blackout_setting,omitempty"`
	UseDiscoveryBasicPollingSettings bool              `json:"use_discovery_basic_polling_settings,omitempty"`
	UseEmailList                     bool              `json:"use_email_list,omitempty"`
	UseEnableDdns                    bool              `json:"use_enable_ddns,omitempty"`
	UseEnableDhcpThresholds          bool              `json:"use_enable_dhcp_thresholds,omitempty"`
	UseEnableDiscovery               bool              `json:"use_enable_discovery,omitempty"`
	UseEnableIfmapPublishing         bool              `json:"use_enable_ifmap_publishing,omitempty"`
	UseIgnoreDhcpOptionListRequest   bool              `json:"use_ignore_dhcp_option_list_request,omitempty"`
	UseIgnoreID                      bool              `json:"use_ignore_id,omitempty"`
	UseIpamEmailAddresses            bool              `json:"use_ipam_email_addresses,omitempty"`
	UseIpamThresholdSettings         bool              `json:"use_ipam_threshold_settings,omitempty"`
	UseIpamTrapSettings              bool              `json:"use_ipam_trap_settings,omitempty"`
	UseLeaseScavengeTime             bool              `json:"use_lease_scavenge_time,omitempty"`
	UseLogicFilterRules              bool              `json:"use_logic_filter_rules,omitempty"`
	UseNextserver                    bool              `json:"use_nextserver,omitempty"`
	UseOptions                       bool              `json:"use_options,omitempty"`
	UsePxeLeaseTime                  bool              `json:"use_pxe_lease_time,omitempty"`
	UseRecycleLeases                 bool              `json:"use_recycle_leases,omitempty"`
	UseSubscribeSettings             bool              `json:"use_subscribe_settings,omitempty"`
	UseUpdateDNSOnLeaseRenewal       bool              `json:"use_update_dns_on_lease_renewal,omitempty"`
	UseZoneAssociations              bool              `json:"use_zone_associations,omitempty"`
	ZoneAssociations                 []ZoneAssociation `json:"zone_associations,omitempty"`
}
