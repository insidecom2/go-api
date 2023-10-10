package response

type LoginAuthRes struct {
	TOKEN 	string `json:"token"`
	REFRESH_TOKEN string `json:"refresh_token"`
}