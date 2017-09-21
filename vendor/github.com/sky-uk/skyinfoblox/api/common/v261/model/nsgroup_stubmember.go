package model

// NSGroupStub : Name Server Group stub object type
type NSGroupStub struct {
	Reference   string         `json:"_ref,omitempty"`
	Name        string         `json:"name,omitempty"`
	Comment     string         `json:"comment,omitempty"`
	StubMembers []MemberServer `json:"stub_members,omitempty"`
}
