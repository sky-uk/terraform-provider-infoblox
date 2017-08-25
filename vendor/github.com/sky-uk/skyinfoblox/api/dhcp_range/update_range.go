package dhcprange

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// UpdateDHCPRangeAPI base object.
type UpdateDHCPRangeAPI struct {
	*api.BaseAPI
}

// NewUpdateDHCPRange updates an existing object
func NewUpdateDHCPRange(dhcpRange DHCPRange) *UpdateDHCPRangeAPI {
	this := new(UpdateDHCPRangeAPI)
	updateEndpoint := fmt.Sprintf("%s/%s", wapiVersion, dhcpRange.Ref)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, updateEndpoint, dhcpRange, new(string))
	return this
}

// GetResponse casts the response object to string
func (ga UpdateDHCPRangeAPI) GetResponse() string {
	return *ga.ResponseObject().(*string)
}
