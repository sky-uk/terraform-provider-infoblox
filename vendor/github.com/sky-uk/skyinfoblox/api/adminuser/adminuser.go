package adminuser

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

var endPoint string

const adminUserEndpoint = "/wapi/v2.6.1"

//NewCreateAdminUser - Create function
func NewCreateAdminUser(newUser AdminUser) *api.BaseAPI {
	endPoint = fmt.Sprintf("%s/adminuser", adminUserEndpoint)
	createUserAPI := api.NewBaseAPI(http.MethodPost, endPoint, newUser, new(string))
	return createUserAPI
}

//NewGetAdminUser - Get a User
func NewGetAdminUser(ref string, returnFields []string) *api.BaseAPI {
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("%s/%s/?_return_fields=%s", adminUserEndpoint, ref, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("%s/%s", adminUserEndpoint, ref)
	}
	updateUserAPI := api.NewBaseAPI(http.MethodGet, endPoint, nil, new(AdminUser))
	return updateUserAPI
}

//NewDeleteAdminUser - Deletes the user
func NewDeleteAdminUser(ref string) *api.BaseAPI {
	endPoint = fmt.Sprintf("%s/%s", adminUserEndpoint, ref)
	deleteUserAPI := api.NewBaseAPI(http.MethodDelete, endPoint, nil, new(string))
	return deleteUserAPI
}

// NewUpdateAdminUser - Updates the user
func NewUpdateAdminUser(updateUser AdminUser) *api.BaseAPI {
	endPoint = fmt.Sprintf("%s/%s", adminUserEndpoint, updateUser.Ref)
	updateUserAPI := api.NewBaseAPI(http.MethodPut, endPoint, updateUser, new(string))
	return updateUserAPI
}
