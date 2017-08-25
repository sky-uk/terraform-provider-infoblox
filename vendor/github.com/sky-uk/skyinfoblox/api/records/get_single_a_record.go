package records

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetSingleARecordAPI base object.
type GetSingleARecordAPI struct {
	*api.BaseAPI
}

// NewGetARecord returns a new object of GetSingleARecordAPI.
func NewGetARecord(recordReference string, returnFields []string) *GetSingleARecordAPI {
	if returnFields != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFields, ",")
		recordReference += returnFields
	}
	this := new(GetSingleARecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, fmt.Sprintf("%s/%s", wapiVersion, recordReference), nil, new(ARecord))
	return this
}

// GetResponse returns ResponseObject of GetSingleARecordAPI.
func (gs GetSingleARecordAPI) GetResponse() ARecord {
	return *gs.ResponseObject().(*ARecord)
}
