package recordusers

import (
	"SignInHelper/pkg/common"
	"SignInHelper/pkg/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *RecordService) AuthCheck(c *gin.Context) {
	res := &common.Response{}
	defer common.WriteResponseWithRecovery(c, res)

	code := c.Query("code")
	fmt.Print(code)
	// get access token
	accessToken, err := rpc.GetAccessToken(code)
	if err != nil {
		http.Redirect(c.Writer, c.Request, "/user.html", http.StatusMovedPermanently)
		common.WriteResponse(res, "get access token failed", err)
		return
	}

	// get user data
	username, err := rpc.GetUserName(accessToken)
	if err != nil {
		http.Redirect(c.Writer, c.Request, "/user.html", http.StatusMovedPermanently)
		common.WriteResponse(res, "get username failed", err)
		return
	}

	// star
	err = rpc.StarProject(accessToken)
	if err != nil {
		http.Redirect(c.Writer, c.Request, "/user.html", http.StatusMovedPermanently)
		common.WriteResponse(res, "star project failed", err)
		return
	}

	// record it in mysql
	err = s.dbClient.InsertNewGithubUser(username, accessToken)
	if err != nil {
		http.Redirect(c.Writer, c.Request, "/user.html", http.StatusMovedPermanently)
		common.WriteResponse(res, "update mysql failed", err)
		return
	}

	// redirect it to form page
	http.Redirect(c.Writer, c.Request, "/signup.html?user="+username, http.StatusMovedPermanently)
}
