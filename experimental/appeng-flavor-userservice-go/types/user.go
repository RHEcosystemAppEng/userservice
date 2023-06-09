package types

type User struct {
	Id string `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Email string `json:"email,omitempty"`

	Emailverified bool `json:"emailVerified,omitempty"`

	Access map[string]bool `json:"access,omitempty"`
}
