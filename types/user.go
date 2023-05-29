package types

type User struct {
	UserId string `json:"id"`

	Username string `json:"username"`

	Email string `json:"email"`

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	OrgAdmin bool `json:"org_admin"`

	IsInternal bool `json:"is_internal"`

	OrgId string `json:"org_id"`

	Type_ string `json:"type"`

	// Added in addition to OpenAPI spec
	Attributes map[string][]string `json:"attributes"`
}
