package dhcprange

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetDHCPRangeAPI base object.
type GetDHCPRangeAPI struct {
	*api.BaseAPI
}

// NewGetDHCPRangeAPI returns a new object of type GetNetworkAPI.
func NewGetDHCPRangeAPI(objRef string, returnFields []string) *GetDHCPRangeAPI {
	if returnFields != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFields, ",")
		objRef += returnFields
	}
	this := new(GetDHCPRangeAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+objRef, nil, new(DHCPRange))
	return this
}

// GetResponse casts the response object and
// returns the single network object
func (gn GetDHCPRangeAPI) GetResponse() DHCPRange {
	return *gn.ResponseObject().(*DHCPRange)
}
