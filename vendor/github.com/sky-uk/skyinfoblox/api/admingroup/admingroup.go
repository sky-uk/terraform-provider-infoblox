package admingroup

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new admin group
func NewCreate(admingroup IBXAdminGroup) *api.BaseAPI {
	createAdminGroupAPI := api.NewBaseAPI(http.MethodPost, adminGroupEndpoint+"/admingroup", admingroup, new(string))
	return createAdminGroupAPI
}

// NewGetAll : used to get all admin groups
func NewGetAll() *api.BaseAPI {
	getAllAdminGroupAPI := api.NewBaseAPI(http.MethodGet, adminGroupEndpoint+"/admingroup", nil, new([]IBXAdminGroupReference))
	return getAllAdminGroupAPI
}

// NewGet : used to get an admin group
func NewGet(ref string, returnFieldList []string) *api.BaseAPI {

	if returnFieldList != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFieldList, ",")
		ref += returnFields
	}
	getAdminGroupAPI := api.NewBaseAPI(http.MethodGet, adminGroupEndpoint+"/"+ref, nil, new(IBXAdminGroup))
	return getAdminGroupAPI
}

// NewUpdate : used to update an admin group
func NewUpdate(adminGroup IBXAdminGroup, returnFields []string) *api.BaseAPI {

	var reference string
	if returnFields != nil {
		reference = "/" + adminGroup.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	} else {
		reference = "/" + adminGroup.Reference
	}
	updateAdminGroupAPI := api.NewBaseAPI(http.MethodPut, adminGroupEndpoint+reference, adminGroup, new(IBXAdminGroup))
	return updateAdminGroupAPI
}

// NewDelete : used to delete an admin group
func NewDelete(reference string) *api.BaseAPI {
	deleteAdminGroupAPI := api.NewBaseAPI(http.MethodDelete, adminGroupEndpoint+"/"+reference, nil, new(string))
	return deleteAdminGroupAPI
}
