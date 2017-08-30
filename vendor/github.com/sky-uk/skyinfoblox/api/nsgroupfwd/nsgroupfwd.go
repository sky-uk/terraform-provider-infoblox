package nsgroupfwd

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new NSGroupFWD object
func NewCreate(nameServerGroupFwd NSGroupFwd) *api.BaseAPI {
	createNSGroupFwdAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+nsGroupFwdEndpoint, nameServerGroupFwd, new(string))
	return createNSGroupFwdAPI
}

// NewGetAll : used to get a list of all NSGroupFWD objects
func NewGetAll() *api.BaseAPI {
	getAllNSGroupFwdAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+nsGroupFwdEndpoint, nil, new([]NSGroupFwd))
	return getAllNSGroupFwdAPI
}

// NewGet : used to get a NSGroupFWD object
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getNSGroupFwdAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+"/"+reference, nil, new(NSGroupFwd))
	return getNSGroupFwdAPI
}

// NewUpdate : used to update a NSGroupFWD object
func NewUpdate(nameServerGroupFwd NSGroupFwd, returnFields []string) *api.BaseAPI {
	reference := "/" + nameServerGroupFwd.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	updateNSGroupFwdAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+reference, nameServerGroupFwd, new(NSGroupFwd))
	return updateNSGroupFwdAPI
}

// NewDelete : used to delete a NSGroupFWD object
func NewDelete(reference string) *api.BaseAPI {
	deleteNSGroupFwdAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+"/"+reference, nil, new(string))
	return deleteNSGroupFwdAPI
}
