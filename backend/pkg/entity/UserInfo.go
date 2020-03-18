package entity

type User struct {
	ID             int64  `json:"id"`
	SSOUsername    string `json:"sso_username"`
	SSOPassword    string `json:"sso_password"`
	Realname       string `json:"realname"`
	SchoolID       string `json:"school_id"`
	Phone          string `json:"phone"`
	GithubUsername string `json:"github_username"`
	GithubToken    string `json:"github_token"`
	Position       string `json:"position"`
	Email          string `json:"email"`
	Count          int64  `json:"count"`
	Stared         int    `json:"stared"`
	Info           string `json:"info"`
}
