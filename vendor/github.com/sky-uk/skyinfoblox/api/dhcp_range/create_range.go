package dhcprange

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// CreateDHCPRangeAPI base object.
type CreateDHCPRangeAPI struct {
	*api.BaseAPI
}

// NewCreateDHCPRange returns a new object of type network.API.
func NewCreateDHCPRange(dhcpRange DHCPRange) *CreateDHCPRangeAPI {
	this := new(CreateDHCPRangeAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/wapi/v2.3.1/range", dhcpRange, new(string))
	return this
}

// GetResponse casts the response object to string
func (ga CreateDHCPRangeAPI) GetResponse() string {
	return *ga.ResponseObject().(*string)
}
