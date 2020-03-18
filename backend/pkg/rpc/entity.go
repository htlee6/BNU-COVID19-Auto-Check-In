package rpc

type getAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type getAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type getUserInfoResponse struct {
	Username string `json:"login"`
}
