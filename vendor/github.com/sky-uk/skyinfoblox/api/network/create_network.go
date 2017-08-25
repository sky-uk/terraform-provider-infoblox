package network

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// CreateNetworkAPI base object.
type CreateNetworkAPI struct {
	*api.BaseAPI
}

// NewCreateNetwork returns a new object of type network.API.
func NewCreateNetwork(net Network) *CreateNetworkAPI {
	this := new(CreateNetworkAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, fmt.Sprintf("%s/network", wapiVersion), net, new(string))
	return this
}

// GetResponse casts the response object to string
func (ga CreateNetworkAPI) GetResponse() string {
	return *ga.ResponseObject().(*string)
}
