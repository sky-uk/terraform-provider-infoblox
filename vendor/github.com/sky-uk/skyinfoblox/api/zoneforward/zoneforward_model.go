package zoneforward

import (
	"github.com/sky-uk/skyinfoblox/api/common"
)

//WapiVersion : WAPI version related with this data model
const WapiVersion = "/wapi/v2.6.1"

// Endpoint - resource WAPI endpoint
const Endpoint = "/zone_forward"

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
	ForwardTo []common.ExternalServer `json:"forward_to,omitempty"`
	//Determines if the appliance sends queries to forwarders only,
	//and not to other internal or Internet root servers.
	ForwardersOnly bool `json:"forwarders_only"`
	//The information for the Grid members to which you want the Infoblox appliance
	//to forward queries for a specified domain name.
	ForwardingServers []common.ForwardingMemberServer `json:"forwarding_servers,omitempty"`
	//The name of this DNS zone. For a reverse zone, this is in “address/cidr” format.
	//For other zones, this is in FQDN format. This value can be in unicode format.
	//Note that for a reverse zone, the corresponding zone_format value should be set.
	//Required
	Fqdn string `json:"fqdn,omitempty"`
	//If you enable this flag, other administrators cannot make conflicting changes.
	//This is for administration purposes only.
	//The zone will continue to serve DNS data even when it is locked.
	Locked bool `json:"locked"`
	//The name of a superuser or the administrator who locked this zone.
	LockedBy string `json:"locked_by,omitempty"`
	//IPv4 Netmask or IPv6 prefix for this zone.
	MaskPrefix string `json:"mask_prefix,omitempty"`
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
	UsingSrgAssociations bool `json:"using_srg_associations,omitempty"`
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
