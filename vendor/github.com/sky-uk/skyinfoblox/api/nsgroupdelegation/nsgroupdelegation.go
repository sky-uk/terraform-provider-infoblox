package nsgroupdelegation

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new NSGroupDelegation object
func NewCreate(nameServerGroupDelegation NSGroupDelegation) *api.BaseAPI {
	createNSGroupDelegationAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+nsGroupDelegationEndpoint, nameServerGroupDelegation, new(string))
	return createNSGroupDelegationAPI
}

// NewGetAll : used to get a list of all NSGroupDelegation objects
func NewGetAll() *api.BaseAPI {
	getAllNSGroupDelegationAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+nsGroupDelegationEndpoint, nil, new([]NSGroupDelegation))
	return getAllNSGroupDelegationAPI
}

// NewGet : used to get a NSGroupDelegation object
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getNSGroupDelegationAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+"/"+reference, nil, new(NSGroupDelegation))
	return getNSGroupDelegationAPI
}

// NewUpdate : used to update a NSGroupDelegation object
func NewUpdate(nameServerGroupDelegation NSGroupDelegation, returnFields []string) *api.BaseAPI {
	reference := "/" + nameServerGroupDelegation.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	updateNSGroupDelegationAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+reference, nameServerGroupDelegation, new(NSGroupDelegation))
	return updateNSGroupDelegationAPI
}

// NewDelete : used to delete a NSGroupDelegation object
func NewDelete(reference string) *api.BaseAPI {
	deleteNSGroupDelegationAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+"/"+reference, nil, new(string))
	return deleteNSGroupDelegationAPI
}
