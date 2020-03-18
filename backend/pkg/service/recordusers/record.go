package recordusers

import (
	"SignInHelper/pkg/common"
	"SignInHelper/pkg/entity"
	"errors"
	"github.com/gin-gonic/gin"
)

type recordRequest struct {
	SSOUsername    string `json:"sso_username"`
	SSOPassword    string `json:"sso_password"`
	Email          string `json:"email"`
	GithubUsername string `json:"github_username"`
	Phone          string `json:"phone"`
}

func (s *RecordService) RecordUser(c *gin.Context) {
	res := &common.Response{}
	defer common.WriteResponseWithRecovery(c, res)

	req := &recordRequest{}
	req.SSOUsername = c.DefaultQuery("sso_username", "")
	req.SSOPassword = c.DefaultQuery("sso_password", "")
	req.Email = c.DefaultQuery("email", "")
	req.GithubUsername = c.DefaultQuery("github_username", "")
	req.Phone = c.DefaultQuery("phone", "")

	user := entity.User{
		SSOUsername:    req.SSOUsername,
		SSOPassword:    req.SSOPassword,
		Email:          req.Email,
		GithubUsername: req.GithubUsername,
		Phone:          req.Phone,
	}

	// check github username exist
	exist, err := s.dbClient.CheckGithubUsernameExist(req.GithubUsername)
	if !exist {
		common.WriteResponse(res, "github not auth", errors.New("please auth github"))
		return
	}

	err = s.dbClient.AddSSOInfo(user)
	if err != nil {
		common.WriteResponse(res, "add user info failed", err)
		return
	}
}
