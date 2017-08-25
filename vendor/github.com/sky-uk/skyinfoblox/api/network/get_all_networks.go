package network

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllNetworksAPI base object.
type GetAllNetworksAPI struct {
	*api.BaseAPI
}

// NewGetAllNetworks returns a new object of GetAllARecordsAPI.
func NewGetAllNetworks(fields []string) *GetAllNetworksAPI {
	this := new(GetAllNetworksAPI)
	var url string
	if len(fields) > 0 {
		url = fmt.Sprintf("%s/network?_return_fields=%s", wapiVersion, strings.Join(fields, ","))
	} else {
		url = fmt.Sprintf("%s/network", wapiVersion)
	}
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new([]Network))
	return this
}

// GetResponse casts the response object and
// returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllNetworksAPI) GetResponse() []Network {
	return *ga.ResponseObject().(*[]Network)
}
