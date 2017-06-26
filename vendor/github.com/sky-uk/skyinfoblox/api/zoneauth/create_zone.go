package zoneauth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// CreateZoneAuthAPI : Create zone API
type CreateZoneAuthAPI struct {
	*api.BaseAPI
}

// NewCreate : Create a new zone
func NewCreate(newZone DNSZone) *CreateZoneAuthAPI {
	this := new(CreateZoneAuthAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/wapi/v2.3.1/zone_auth", newZone, new(string))
	return this
}

// GetResponse : get response object from created zone
func (cza *CreateZoneAuthAPI) GetResponse() string {
	return *cza.ResponseObject().(*string)
}
