package zoneauth

import (
	"fmt"
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
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, fmt.Sprintf("%s/zone_auth", wapiVersion), newZone, new(string))
	return this
}

// GetResponse : get response object from created zone
func (cza *CreateZoneAuthAPI) GetResponse() string {
	return *cza.ResponseObject().(*string)
}
