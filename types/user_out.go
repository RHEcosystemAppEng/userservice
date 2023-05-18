package types

import (
	"time"
)

type UserOut struct {
	Uuid string `json:"uuid"`

	Created time.Time `json:"created"`

	Modified time.Time `json:"modified"`

	UserId string `json:"user_id,omitempty"`

	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	OrgAdmin bool `json:"org_admin,omitempty"`

	IsInternal bool `json:"is_internal,omitempty"`

	OrgId string `json:"org_id,omitempty"`

	Type_ string `json:"type,omitempty"`

	// Added in addition to OpenAPI spec
	Attributes map[string][]string `json:"attributes,omitempty"`
}
