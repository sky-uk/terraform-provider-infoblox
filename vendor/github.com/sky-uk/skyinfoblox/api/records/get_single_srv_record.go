package records

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetSingleSRVRecordAPI base object.
type GetSingleSRVRecordAPI struct {
	*api.BaseAPI
}

// NewGetSRVRecord returns a new object of GetSingleSRVRecordAPI.
func NewGetSRVRecord(recordReference string, returnFields []string) *GetSingleSRVRecordAPI {
	if returnFields != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFields, ",")
		recordReference += returnFields
	}
	this := new(GetSingleSRVRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, fmt.Sprintf("%s/%s", wapiVersion, recordReference), nil, new(SRVRecord))
	return this
}

// GetResponse returns ResponseObject of GetSingleSRVRecordAPI.
func (gs GetSingleSRVRecordAPI) GetResponse() SRVRecord {
	return *gs.ResponseObject().(*SRVRecord)
}
