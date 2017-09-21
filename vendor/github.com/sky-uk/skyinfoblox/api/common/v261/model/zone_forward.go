package model

// ZoneForward - the zone forward resource data model
type ZoneForward struct {
	// The object reference. Cannot be set.
	Ref string `json:"_ref,omitempty"`
	//The IP address of the server that is serving this zone.
	//Not writable
	Address string `json:"address,omitempty"`
	//Comment for the zone; maximum 256 characters.
	Comment string `json:"comment,omitempty"`
	//Determines whether a zone is disabled or not. When this is set to False, the zone is enabled.
	Disable bool `json:"disable"`
	//The displayed name of the DNS zone.
	//Not writable
	DisplayDomain string `json:"display_domain,omitempty"`
	//The name of this DNS zone in punycode format.
	//For a reverse zone, this is in “address/cidr” format.
	//For other zones, this is in FQDN format in punycode format.
	DNSFqdn string `json:"dns_fqdn,omitempty"`
	//Extensible attributes associated with the object.
	ExtAttrs string `json:"extattrs,omitempty"`
	//The information for the remote name servers to which you want the Infoblox appliance
	//to forward queries for a specified domain name.
	//Required
	ForwardTo []ExternalServer `json:"forward_to,omitempty"`
	//Determines if the appliance sends queries to forwarders only,
	//and not to other internal or Internet root servers.
	ForwardersOnly bool `json:"forwarders_only"`
	//The information for the Grid members to which you want the Infoblox appliance
	//to forward queries for a specified domain name.
	ForwardingServers []ForwardingMemberServer `json:"forwarding_servers,omitempty"`
	//The name of this DNS zone. For a reverse zone, this is in “address/cidr” format.
	//For other zones, this is in FQDN format. This value can be in unicode format.
	//Note that for a reverse zone, the corresponding zone_format value should be set.
	//Required
	// This field cannot be changed at update , should be omitempty to avoid sending empty string
	Fqdn string `json:"fqdn,omitempty"`
	//If you enable this flag, other administrators cannot make conflicting changes.
	//This is for administration purposes only.
	//The zone will continue to serve DNS data even when it is locked.
	Locked bool `json:"locked"`
	//The name of a superuser or the administrator who locked this zone.
	LockedBy string `json:"locked_by,omitempty"`
	//IPv4 Netmask or IPv6 prefix for this zone.
	MaskPrefix string `json:"mask_prefix,omitempty"`
	//The flag that determines whether Active Directory is integrated or not.
	//This field is valid only when ms_managed is “STUB”, “AUTH_PRIMARY”, or “AUTH_BOTH”.
	MSADIntegrated bool `json:"ms_ad_integrated"`
	//Determines whether an Active Directory-integrated zone with a Microsoft DNS server
	//as primary allows dynamic updates. Valid values are:
	//  “SECURE” if the zone allows secure updates only.
	//  “NONE” if the zone forbids dynamic updates.
	//  “ANY” if the zone accepts both secure and nonsecure updates.
	//This field is valid only if ms_managed is either “AUTH_PRIMARY” or “AUTH_BOTH”.
	//If the flag ms_ad_integrated is false, the value “SECURE” is not allowed.
	MSDDNSMode string `json:"ms_ddns_mode,omitempty"`
	//The flag that indicates whether the zone is assigned to a Microsoft DNS server.
	//This flag returns the authoritative name server type of the Microsoft DNS server.
	//Valid values are:
	// “NONE” if the zone is not assigned to any Microsoft DNS server.
	// “STUB” if the zone is assigned to a Microsoft DNS server as a stub zone.
	// “AUTH_PRIMARY” if only the primary server of the zone is a Microsoft DNS server.
	// “AUTH_SECONDARY” if only the secondary server of the zone is a Microsoft DNS server.
	// “AUTH_BOTH” if both the primary and secondary servers of the zone are Microsoft DNS servers.
	MSManaged string `json:"ms_managed,omitempty"`
	//Determines if a Grid member manages the zone served by a Microsoft DNS server in read-only mode.
	//This flag is true when a Grid member manages the zone in read-only mode, false otherwise.
	//When the zone has the ms_read_only flag set to True, no changes can be made to this zone.
	//MSReadOnly bool `json:"ms_read_only,omitempty"`
	//The name of MS synchronization master for this zone.
	MSSyncMasterName string `json:"ms_sync_master_name,omitempty"`
	//A forwarding member name server group. Values with leading or trailing white space are not valid for this field.
	//The default value is undefined.
	NSGroup string `json:"ns_group,omitempty"`
	//The parent zone of this zone.
	//Note that when searching for reverse zones, the “in-addr.arpa” notation should be used.
	// Not writable
	Parent string `json:"parent,omitempty"`
	//The RFC2317 prefix value of this DNS zone.
	//Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits.
	//Enter a prefix, such as the name of the allocated address block.
	//The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B.
	Prefix string `json:"prefix,omitempty"`
	//This is true if the zone is associated with a shared record group.
	//UsingSrgAssociations bool `json:"using_srg_associations,omitempty"`
	//The name of the DNS view in which the zone resides. Example “external”.
	View string `json:"view,omitempty"`
	//Determines the format of this zone.
	//Valid values are:
	//  FORWARD
	//  IPV4
	//  IPV6
	//The default value is FORWARD.
	ZoneFormat string `json:"zone_format,omitempty"`
}
