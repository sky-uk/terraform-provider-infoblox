package common

// ErrorStruct - the error structure received back from a faulty api call
type ErrorStruct struct {
	Error string `json:"Error"`
	Code  string `json:"code"`
	Text  string `json:"text"`
	Trace string `json:"trace"`
}
