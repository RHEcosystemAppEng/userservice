package types

type User struct {
	UserId string `json:"user_id,omitempty"`

	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	OrgAdmin bool `json:"org_admin,omitempty"`

	IsInternal bool `json:"is_internal,omitempty"`

	OrgId string `json:"org_id,omitempty"`

	Type_ string `json:"type,omitempty"`
}
