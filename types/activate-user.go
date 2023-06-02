package types

type ActivateUser struct {
	Enabled    string              `json:"enabled"`
	Attributes map[string][]string `json:"attributes"`
}
