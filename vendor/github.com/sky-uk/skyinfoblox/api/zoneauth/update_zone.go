package zoneauth

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// UpdateZoneAuthAPI : Update zone API
type UpdateZoneAuthAPI struct {
	*api.BaseAPI
}

// NewUpdate : Update zone
func NewUpdate(updateDNSZone DNSZone, returnFieldList []string) *UpdateZoneAuthAPI {

	var reference string

	if returnFieldList != nil {
		reference = updateDNSZone.Reference + "?_return_fields=" + strings.Join(returnFieldList, ",")
	} else {
		reference = updateDNSZone.Reference
	}
	this := new(UpdateZoneAuthAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, fmt.Sprintf("%s/%s", wapiVersion, reference), updateDNSZone, new(DNSZone))
	return this
}

// GetResponse : returns the response from UpdateZoneAPI
func (updateZoneAPI *UpdateZoneAuthAPI) GetResponse() DNSZone {
	return *updateZoneAPI.ResponseObject().(*DNSZone)
}
