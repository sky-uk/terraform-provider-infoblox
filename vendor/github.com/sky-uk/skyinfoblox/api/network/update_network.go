package network

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// UpdateNetworkAPI base object.
type UpdateNetworkAPI struct {
	*api.BaseAPI
}

// NewUpdateNetwork returns a new object of type UpdateNetworkAPI.
func NewUpdateNetwork(updatedObj Network) *UpdateNetworkAPI {
	this := new(UpdateNetworkAPI)
	qPath := fmt.Sprintf("%s/%s", wapiVersion, updatedObj.Ref)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, qPath, updatedObj, new(string))
	return this
}

// GetResponse casts the response object and returns the string representing
// the updated object reference or nil in case of errors
func (ga UpdateNetworkAPI) GetResponse() string {
	return *ga.ResponseObject().(*string)
}
