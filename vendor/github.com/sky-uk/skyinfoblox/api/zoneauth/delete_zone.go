package zoneauth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// DeleteZoneAuthAPI : Zone API for deleting a zone
type DeleteZoneAuthAPI struct {
	*api.BaseAPI
}

// NewDelete : delete a resource by it's reference - this function can probably be common to all Infoblox resource types.
func NewDelete(ref string) *DeleteZoneAuthAPI {
	this := new(DeleteZoneAuthAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/wapi/v2.3.1/"+ref, nil, new(string))
	return this
}

// GetResponse : return response object
func (dza DeleteZoneAuthAPI) GetResponse() string {
	return *dza.ResponseObject().(*string)
}
