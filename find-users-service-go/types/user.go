package types

type User struct {
	UserId string `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	OrgAdmin bool `json:"org_admin,omitempty"`

	IsInternal bool `json:"is_internal,omitempty"`

	OrgId string `json:"org_id,omitempty"`

	Type_ string `json:"type,omitempty"`

	// Added in addition to OpenAPI spec
	Attributes map[string][]string `json:"attributes,omitempty"`
}
