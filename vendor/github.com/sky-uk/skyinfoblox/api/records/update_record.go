package records

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// UpdateRecordAPI base object.
type UpdateRecordAPI struct {
	*api.BaseAPI
}

// NewUpdateRecord returns a new object of UpdateRecordAPI.
func NewUpdateRecord(recordReference string, requestPayload GenericRecord) *UpdateRecordAPI {
	this := new(UpdateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, fmt.Sprintf("%s/%s", wapiVersion, recordReference), requestPayload, new(string))
	return this
}

// NewUpdateARecord returns a new object of UpdateRecordAPI.
func NewUpdateARecord(recordReference string, requestPayload ARecord) *UpdateRecordAPI {
	this := new(UpdateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, fmt.Sprintf("%s/%s", wapiVersion, recordReference), requestPayload, new(string))
	return this
}

// GetResponse returns ResponseObject of UpdateARecordAPI.
func (u UpdateRecordAPI) GetResponse() string {
	return *u.ResponseObject().(*string)
}
