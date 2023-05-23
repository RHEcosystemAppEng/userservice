package types

type FindUsersCriteria struct {
	OrgId               string `form:"org_id"`
	EmailsQueryArray    string `form:"emails"`
	UserIdsQueryArray   string `form:"user_ids"`
	UserNamesQueryArray string `form:"usernames"`
	QueryLimit          int    `form:"limit,default=1" binding:"omitempty,numeric,gte=1,max=1000"` // Max number of users to return
	Offset              int    `form:"offset,default=0" binding:"omitempty,numeric,gte=0"`
	OrderBy             string `form:"order" binding:"omitempty,oneof=email username modified created"` // values from specs: mail, username, modified, created
	OrderDirection      string `form:"direction" binding:"omitempty,oneof=asc desc"`                    // introduced by us, to be used only when Order parameter is specified. values: asc, desc

	Emails []string

	Usernames []string

	UserIds []string
}
