package zonestub

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

var endPoint string

// NewCreate - Create a new sbut zone
func NewCreate(newzoneStub ZoneStub) *api.BaseAPI {
	endPoint = fmt.Sprintf("%s/%s", WapiVersion, Endpoint)
	createAPI := api.NewBaseAPI(http.MethodPost, endPoint, newzoneStub, new(string))
	return createAPI
}

//NewGet - Get a single stub zone
func NewGet(ref string, returnFields []string) *api.BaseAPI {
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("%s/%s?_return_fields=%s", WapiVersion, ref, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("%s/%s", WapiVersion, ref)
	}
	getAPI := api.NewBaseAPI(http.MethodGet, endPoint, nil, new(ZoneStub))
	return getAPI
}

// NewGetAll - Get all stub zones
func NewGetAll(returnFields []string) *api.BaseAPI {
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("%s/%s?_return_fields=%s", WapiVersion, Endpoint, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("%s/%s", WapiVersion, Endpoint)
	}
	getAllAPI := api.NewBaseAPI(http.MethodGet, endPoint, nil, new([]ZoneStub))
	return getAllAPI

}

// NewUpdate - Updates an existing Zone
func NewUpdate(updateZoneStub ZoneStub) *api.BaseAPI {
	endPoint = fmt.Sprintf("%s/%s", WapiVersion, updateZoneStub.Ref)
	updateAPI := api.NewBaseAPI(http.MethodPut, endPoint, updateZoneStub, new(string))
	return updateAPI
}

// NewDelete - Deletes an existing zone
func NewDelete(ref string) *api.BaseAPI {
	endPoint = fmt.Sprintf("%s/%s", WapiVersion, ref)
	deleteAPI := api.NewBaseAPI(http.MethodDelete, endPoint, nil, new(string))
	return deleteAPI
}
