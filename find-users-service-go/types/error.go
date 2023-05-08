package types

type Error struct {
	Detail string `json:"detail,omitempty"`

	Status string `json:"status,omitempty"`
}
