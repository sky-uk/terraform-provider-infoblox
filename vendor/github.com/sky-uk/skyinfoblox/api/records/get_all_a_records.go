package records

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllARecordsAPI base object.
type GetAllARecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllARecords returns a new object of GetAllARecordsAPI.
func NewGetAllARecords(fields []string) *GetAllARecordsAPI {
	var url string
	if len(fields) >= 1 {
		url = fmt.Sprintf("%s/record:a?_return_fields=%s", wapiVersion, strings.Join(fields, ","))
	} else {
		url = fmt.Sprintf("%s/record:a", wapiVersion)
	}

	this := new(GetAllARecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new([]ARecord))
	return this
}

// GetResponse returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllARecordsAPI) GetResponse() []ARecord {
	return *ga.ResponseObject().(*[]ARecord)
}
