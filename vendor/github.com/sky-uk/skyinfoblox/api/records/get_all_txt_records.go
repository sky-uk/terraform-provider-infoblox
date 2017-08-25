package records

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllTXTRecordsAPI base object.
type GetAllTXTRecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllTXTRecords returns a new object of GetAllTXTRecordsAPI.
func NewGetAllTXTRecords(fields []string) *GetAllTXTRecordsAPI {
	var url string
	if len(fields) >= 1 {
		url = fmt.Sprintf("%s/record:txt?_return_fields=%s", wapiVersion, strings.Join(fields, ","))
	} else {
		url = fmt.Sprintf("%s/record:txt", wapiVersion)
	}
	this := new(GetAllTXTRecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new([]TXTRecord))
	return this
}

// GetResponse returns ResponseObject of GetAllTXTRecordsAPI.
func (ga GetAllTXTRecordsAPI) GetResponse() []TXTRecord {
	return *ga.ResponseObject().(*[]TXTRecord)
}
