package types

type PaginationMeta struct {
	Total int64 `json:"total,omitempty"`

	First string `json:"first,omitempty"`

	Previous string `json:"previous,omitempty"`

	Next string `json:"next,omitempty"`

	Last string `json:"last,omitempty"`
}
