package adminrole

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

const adminRoleEndpoint = "/wapi/v2.6.1/"
const returnFields = "?_return_fields=name,comment,disable"

// NewGet : used to get an admin role
func NewGet(roleRef string) *api.BaseAPI {
	getAdminRoleAPI := api.NewBaseAPI(http.MethodGet, adminRoleEndpoint+roleRef+returnFields, nil, new(AdminRole))
	return getAdminRoleAPI
}

// NewGetAll : used to get all admin roles
func NewGetAll() *api.BaseAPI {
	getAllAdminRolesAPI := api.NewBaseAPI(http.MethodGet, adminRoleEndpoint+"adminrole"+returnFields, nil, new([]AdminRole))
	return getAllAdminRolesAPI
}

// NewCreate : used to create a new admin role
func NewCreate(adminRole AdminRole) *api.BaseAPI {
	createAdminRoleAPI := api.NewBaseAPI(http.MethodPost, adminRoleEndpoint+"adminrole", adminRole, new(string))
	return createAdminRoleAPI
}

// NewUpdate : used to update an admin role
func NewUpdate(roleRef string, adminRole AdminRole) *api.BaseAPI {
	updateAdminRoleAPI := api.NewBaseAPI(http.MethodPut, adminRoleEndpoint+roleRef+returnFields, adminRole, new(AdminRole))
	return updateAdminRoleAPI
}

// NewDelete : used to delete an admin role
func NewDelete(roleRef string) *api.BaseAPI {
	deleteAdminRoleAPI := api.NewBaseAPI(http.MethodDelete, adminRoleEndpoint+roleRef, nil, new(string))
	return deleteAdminRoleAPI
}
