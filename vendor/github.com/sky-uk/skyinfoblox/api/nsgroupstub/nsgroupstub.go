package nsgroupstub

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new NSGroupStub object
func NewCreate(nameServerGroupStub NSGroupStub) *api.BaseAPI {
	createNSGroupStubAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+nsGroupStubEndpoint, nameServerGroupStub, new(string))
	return createNSGroupStubAPI
}

// NewGetAll : used to get a list of all NSGroupStub objects
func NewGetAll() *api.BaseAPI {
	getAllNSGroupStubAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+nsGroupStubEndpoint, nil, new([]NSGroupStub))
	return getAllNSGroupStubAPI
}

// NewGet : used to get a NSGroupStub object
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getNSGroupStubAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+"/"+reference, nil, new(NSGroupStub))
	return getNSGroupStubAPI
}

// NewUpdate : used to update a NSGroupStub object
func NewUpdate(nameServerGroupStub NSGroupStub, returnFields []string) *api.BaseAPI {
	reference := "/" + nameServerGroupStub.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	updateNSGroupStubAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+reference, nameServerGroupStub, new(NSGroupStub))
	return updateNSGroupStubAPI
}

// NewDelete : used to delete a NSGroupStub object
func NewDelete(reference string) *api.BaseAPI {
	deleteNSGroupStubAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+"/"+reference, nil, new(string))
	return deleteNSGroupStubAPI
}
