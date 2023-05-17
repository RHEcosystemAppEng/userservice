package types

type FindUsersCriteria struct {
	OrgId               string `form:"org_id"`
	EmailsQueryArray    string `form:"emails"`
	UserIdsQueryArray   string `form:"user_ids"`
	UserNamesQueryArray string `form:"usernames"`
	QueryLimit          int    `form:"limit" binding:"omitempty,numeric,gte=1,max=1000"` // Max number of users to return

	Emails []string

	Usernames []string

	UserIds []string
}
