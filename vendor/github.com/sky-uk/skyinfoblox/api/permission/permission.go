package permission

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

const permissionEndpoint = "/wapi/v2.6.1/"
const returnFields = "?_return_fields=group,object,permission,resource_type,role"

// NewGet returns a new object of permissionGetAPI.
func NewGet(permissionRef string) *api.BaseAPI {
	getPermissionAPI := api.NewBaseAPI(http.MethodGet, permissionEndpoint+permissionRef+returnFields, nil, new(Permission))
	return getPermissionAPI
}

// NewGetAll returns a new object of permissionGetAllAPI.
func NewGetAll() *api.BaseAPI {
	getAllPermissionsAPI := api.NewBaseAPI(http.MethodGet, permissionEndpoint+"permission"+returnFields, nil, new([]Permission))
	return getAllPermissionsAPI
}

// NewCreate returns a new object of permissionCreateAPI.
func NewCreate(newPermission Permission) *api.BaseAPI {
	createPermissionAPI := api.NewBaseAPI(http.MethodPost, permissionEndpoint+"permission", newPermission, new(string))
	return createPermissionAPI
}

// NewUpdate returns a new object of permissionUpdateAPI.
func NewUpdate(permissionRef string, updatedPermission Permission) *api.BaseAPI {
	updatePermissionAPI := api.NewBaseAPI(http.MethodPut, permissionEndpoint+permissionRef+returnFields, updatedPermission, new(Permission))
	return updatePermissionAPI
}

// NewDelete returns a new object of permissionDeleteAPI.
func NewDelete(permissionRef string) *api.BaseAPI {
	deletePermissionAPI := api.NewBaseAPI(http.MethodDelete, permissionEndpoint+permissionRef, nil, new(string))
	return deletePermissionAPI
}
