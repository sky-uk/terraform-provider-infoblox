package records

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// CreateRecordAPI base object.
type CreateRecordAPI struct {
	*api.BaseAPI
}

// NewCreateRecord returns a new object of CreateRecordAPI.
func NewCreateRecord(recordType string, requestPayload GenericRecord) *CreateRecordAPI {
	this := new(CreateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, fmt.Sprintf("%s/record:%s", wapiVersion, recordType), requestPayload, new(string))
	return this
}

// NewCreateARecord - Creates a new A record
func NewCreateARecord(requestPayload ARecord) *CreateRecordAPI {
	this := new(CreateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, fmt.Sprintf("%s/record:a", wapiVersion), requestPayload, new(string))
	return this
}

// NewCreateTXTRecord - Creates a new A record
func NewCreateTXTRecord(requestPayload TXTRecord) *CreateRecordAPI {
	this := new(CreateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, fmt.Sprintf("%s/record:txt", wapiVersion), requestPayload, new(string))
	return this
}

// GetResponse returns ResponseObject of CreateRecordAPI.
func (c CreateRecordAPI) GetResponse() string {
	return *c.ResponseObject().(*string)
}
