package types

type UserPagination struct {
	Meta *PaginationMeta `json:"meta,omitempty"`

	Users []UserOut `json:"users"`
}
