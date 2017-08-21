package zoneforward

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new forward zone
func NewCreate(zoneForward ZoneForward) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodPost, WapiVersion+Endpoint, zoneForward, new(string))
}

// NewGetAll : used to get all admin groups
func NewGetAll() *api.BaseAPI {
	return api.NewBaseAPI(http.MethodGet, WapiVersion+Endpoint, nil, new([]ZoneForward))
}

// NewGet : used to get an admin group
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	if returnFieldList != nil && len(returnFieldList) > 0 {
		reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	}
	return api.NewBaseAPI(http.MethodGet, WapiVersion+"/"+reference, nil, new(ZoneForward))
}

// NewUpdate : used to update an admin group
func NewUpdate(zoneForward ZoneForward, returnFields []string) *api.BaseAPI {
	reference := "/" + zoneForward.Ref + "?_return_fields=" + strings.Join(returnFields, ",")
	return api.NewBaseAPI(http.MethodPut, WapiVersion+reference, zoneForward, new(ZoneForward))
}

// NewDelete : used to delete an admin group
func NewDelete(reference string) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodDelete, WapiVersion+"/"+reference, nil, new(string))
}
