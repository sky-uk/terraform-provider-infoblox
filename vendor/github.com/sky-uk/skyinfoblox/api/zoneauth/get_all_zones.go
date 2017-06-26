package zoneauth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// GetAllZoneAuthAPI : all zones struct
type GetAllZoneAuthAPI struct {
	*api.BaseAPI
}

// NewGetAllZones : returns an object containing all zones.
func NewGetAllZones() *GetAllZoneAuthAPI {
	this := new(GetAllZoneAuthAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/zone_auth?_return_fields=fqdn", nil, new(DNSZoneReferences))
	return this
}

// GetResponse : returns the response object of GetAllZones
func (gaz GetAllZoneAuthAPI) GetResponse() *DNSZoneReferences {
	return gaz.ResponseObject().(*DNSZoneReferences)
}
