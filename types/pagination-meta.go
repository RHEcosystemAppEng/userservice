package types

type PaginationMeta struct {
	Total int64 `json:"total"`

	First string `json:"first"`

	Previous string `json:"previous"`

	Next string `json:"next"`

	Last string `json:"last"`
}
