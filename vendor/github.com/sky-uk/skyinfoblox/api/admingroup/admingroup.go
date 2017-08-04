package admingroup

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new admin group
func NewCreate(admingroup IBXAdminGroup) *api.BaseAPI {
	createAdminGroupAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+adminGroupEndpoint, admingroup, new(string))
	return createAdminGroupAPI
}

// NewGetAll : used to get all admin groups
func NewGetAll() *api.BaseAPI {
	getAllAdminGroupAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+adminGroupEndpoint, nil, new([]IBXAdminGroupReference))
	return getAllAdminGroupAPI
}

// NewGet : used to get an admin group
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getAdminGroupAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+"/"+reference, nil, new(IBXAdminGroup))
	return getAdminGroupAPI
}

// NewUpdate : used to update an admin group
func NewUpdate(adminGroup IBXAdminGroup, returnFields []string) *api.BaseAPI {
	reference := "/" + adminGroup.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	updateAdminGroupAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+reference, adminGroup, new(IBXAdminGroup))
	return updateAdminGroupAPI
}

// NewDelete : used to delete an admin group
func NewDelete(reference string) *api.BaseAPI {
	deleteAdminGroupAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+"/"+reference, nil, new(string))
	return deleteAdminGroupAPI
}
