package nsgroupfwdstub

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new NSGroupFwdStub object
func NewCreate(nameServerGroupFwdStub NSGroupFwdStub) *api.BaseAPI {
	createNSGroupFwdStubAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+nsGroupFwdStubEndpoint, nameServerGroupFwdStub, new(string))
	return createNSGroupFwdStubAPI
}

// NewGetAll : used to get a list of all NSGroupFwdStub objects
func NewGetAll() *api.BaseAPI {
	getAllNSGroupFwdStubAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+nsGroupFwdStubEndpoint, nil, new([]NSGroupFwdStub))
	return getAllNSGroupFwdStubAPI
}

// NewGet : used to get a NSGroupFwdStub object
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getNSGroupFwdStubAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+"/"+reference, nil, new(NSGroupFwdStub))
	return getNSGroupFwdStubAPI
}

// NewUpdate : used to update a NSGroupFwdStub object
func NewUpdate(nameServerGroupFwdStub NSGroupFwdStub, returnFields []string) *api.BaseAPI {
	reference := "/" + nameServerGroupFwdStub.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	updateNSGroupFwdStubAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+reference, nameServerGroupFwdStub, new(NSGroupFwdStub))
	return updateNSGroupFwdStubAPI
}

// NewDelete : used to delete a NSGroupFwdStub object
func NewDelete(reference string) *api.BaseAPI {
	deleteNSGroupFwdStubAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+"/"+reference, nil, new(string))
	return deleteNSGroupFwdStubAPI
}
