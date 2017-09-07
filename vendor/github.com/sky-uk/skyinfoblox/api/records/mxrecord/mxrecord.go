package mxrecord

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate - Creates a new Record
func NewCreate(mxRecord MxRecord) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodPost, fmt.Sprintf("%s/%s", wapiVersion, mxRecordEndpoint), mxRecord, new(string))
}

// NewGet - Returns a single record
func NewGet(reference string, returnFields []string) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodGet, fmt.Sprintf("%s/%s?_return_fields=%s", wapiVersion, reference, strings.Join(returnFields, ",")), nil, new(MxRecord))
}

// NewGetAll - Returns all records
func NewGetAll() *api.BaseAPI {
	return api.NewBaseAPI(http.MethodGet, fmt.Sprintf("%s/%s", wapiVersion, mxRecordEndpoint), nil, new([]MxRecord))

}

// NewUpdate - Updates a Record
func NewUpdate(reference string, mxRecord MxRecord) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodPut, fmt.Sprintf("%s/%s", wapiVersion, reference), mxRecord, new(string))
}

// NewDelete - Deletes a Record
func NewDelete(reference string) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodDelete, fmt.Sprintf("%s/%s", wapiVersion, reference), nil, new(string))

}
