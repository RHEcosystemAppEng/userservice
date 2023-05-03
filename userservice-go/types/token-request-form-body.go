package types

type TokenRequestFormBody struct {
	Username   string `form:"username"`
	Password   string `form:"password"`
	Grant_type string `form:"grant_type"`
	Client_id  string `form:"client_id"`
}
