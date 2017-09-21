package rest

// BaseAPI  - Base API struct.
type BaseAPI struct {
	method         string
	endpoint       string
	requestObject  interface{}
	responseObject interface{}
	errorObject    interface{}
	statusCode     int
	rawResponse    []byte
	err            error
}

// NewBaseAPI - Returns a new object of the BaseAPI.
func NewBaseAPI(
	method string,
	endpoint string,
	requestObject interface{},
	responseObject interface{},
	errorObject interface{},
) *BaseAPI {
	return &BaseAPI{method, endpoint, requestObject, responseObject, errorObject, 0, nil, nil}
}

// RequestObject - Returns the request object of the BaseAPI
func (b *BaseAPI) RequestObject() interface{} {
	return b.requestObject
}

// ResponseObject - Returns the ResponseObject interface.
func (b *BaseAPI) ResponseObject() interface{} {
	return b.responseObject
}

// ErrorObject - Returns the ErrorObject interface.
func (b *BaseAPI) ErrorObject() interface{} {
	return b.errorObject
}

// Method - Returns the Method string, i.e. GET, PUT, POST.
func (b *BaseAPI) Method() string {
	return b.method
}

// Endpoint - Returns the Endpoint url string.
func (b *BaseAPI) Endpoint() string {
	return b.endpoint
}

// StatusCode - Returns the status code of the api.
func (b *BaseAPI) StatusCode() int {
	return b.statusCode
}

// RawResponse - Returns the rawResponse object as byte type.
func (b *BaseAPI) RawResponse() []byte {
	return b.rawResponse
}

// Error - Returns the err the api.
func (b *BaseAPI) Error() error {
	return b.err
}

// SetStatusCode - Sets the statusCode from api object.
func (b *BaseAPI) SetStatusCode(statusCode int) {
	b.statusCode = statusCode
}

// SetRawResponse - Sets the rawResponse on api object.
func (b *BaseAPI) SetRawResponse(rawResponse []byte) {
	b.rawResponse = rawResponse
}

// SetError - Sets the err on api object.
func (b *BaseAPI) SetError(err error) {
	b.err = err
}

// SetResponseObject - Sets the responseObject
func (b *BaseAPI) SetResponseObject(res interface{}) {
	b.responseObject = res
}

// SetErrorObject - Sets the errorObject
func (b *BaseAPI) SetErrorObject(res interface{}) {
	b.errorObject = res
}
