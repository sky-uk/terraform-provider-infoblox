package zonedelegated

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

var endPoint string

// NewCreate - Create a new Zone
func NewCreate(newZoneDelegated ZoneDelegated) *api.BaseAPI {
	endPoint = fmt.Sprintf("%s/zone_delegated", wapiVersion)
	createZoneAPI := api.NewBaseAPI(http.MethodPost, endPoint, newZoneDelegated, new(string))
	return createZoneAPI
}

// NewGet - Read an existing zone
func NewGet(ref string, returnFields []string) *api.BaseAPI {
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("%s/%s/?_return_fields=%s", wapiVersion, ref, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("%s/%s", wapiVersion, ref)
	}
	getZoneAPI := api.NewBaseAPI(http.MethodGet, endPoint, nil, new(ZoneDelegated))
	return getZoneAPI

}

// NewGetAll - Get all existing zones
func NewGetAll(returnFields []string) *api.BaseAPI {
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("%s/zone_delegated?_return_fields=%s", wapiVersion, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("%s/zone_delegated", wapiVersion)
	}
	getAllZoneAPI := api.NewBaseAPI(http.MethodGet, endPoint, nil, new([]ZoneDelegated))
	return getAllZoneAPI
}

// NewUpdate - Update a zone
func NewUpdate(ref string, updateZoneDelegated ZoneDelegated) *api.BaseAPI {
	endPoint := fmt.Sprintf("%s/%s", wapiVersion, ref)
	updateZoneAPI := api.NewBaseAPI(http.MethodPut, endPoint, updateZoneDelegated, new(string))
	return updateZoneAPI

}

// NewDelete - Delete a zone
func NewDelete(ref string) *api.BaseAPI {
	endPoint := fmt.Sprintf("%s/%s", wapiVersion, ref)
	deleteZoneAPI := api.NewBaseAPI(http.MethodDelete, endPoint, nil, new(string))
	return deleteZoneAPI

}
