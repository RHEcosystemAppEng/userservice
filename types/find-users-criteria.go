package types

type FindUsersCriteria struct {
	OrgId               string `form:"org_id"`
	EmailsQueryArray    string `form:"emails"`
	UserIdsQueryArray   string `form:"user_ids"`
	UserNamesQueryArray string `form:"usernames"`
	QueryLimit          int    `form:"limit,default=1" binding:"omitempty,numeric,min=1,max=1000"` // Max number of users to return
	Offset              int    `form:"offset,default=0" binding:"omitempty"`

	Emails []string

	Usernames []string

	UserIds []string
}
