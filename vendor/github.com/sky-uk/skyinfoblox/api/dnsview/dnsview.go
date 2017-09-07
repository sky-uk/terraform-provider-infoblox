package dnsview

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

//NewCreate : used to create a new DNSView object
func NewCreate(dnsView DNSView) *api.BaseAPI {
	createDNSViewAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+dnsViewEndpoint, dnsView, new(string))
	return createDNSViewAPI
}

//NewUpdate : used to update an existing DNSView object
func NewUpdate(dnsView DNSView, returnFieldList []string) *api.BaseAPI {
	reference := dnsView.Reference + "?_return_fields=" + strings.Join(returnFieldList, ",")
	createDNSViewAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+reference, dnsView, new(DNSView))
	return createDNSViewAPI
}

//NewGetAll : used to retrieve all DNSView objects
func NewGetAll() *api.BaseAPI {
	getAllDNSViewAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+dnsViewEndpoint, nil, new([]DNSView))
	return getAllDNSViewAPI
}

//NewGet : used to retrieve a DNSView object
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getDNSViewAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+reference, nil, new(DNSView))
	return getDNSViewAPI
}

//NewDelete : used to delete a DNSView object
func NewDelete(reference string) *api.BaseAPI {
	deleteDNSViewAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+reference, nil, new(string))
	return deleteDNSViewAPI
}
