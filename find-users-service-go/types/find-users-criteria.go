package types

type FindUsersCriteria struct {
	OrgId               string `form:"org_id" binding:"required"`
	EmailsQueryArray    string `form:"emails"`
	UserIdsQueryArray   string `form:"user_ids"`
	UserNamesQueryArray string `form:"usernames"`

	Emails []string

	Usernames []string

	UserIds []string
}
