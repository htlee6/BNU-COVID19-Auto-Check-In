package common

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}

func WriteResponseWithRecovery(c *gin.Context, res *Response) {
	if err := recover(); err != nil {
		res.Code = http.StatusBadRequest
		res.Error = Error{Message: "interface exception!"}
		c.JSON(http.StatusBadRequest, res)
	}
	if res.Code == http.StatusBadRequest {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Code = http.StatusOK
	res.Message = "success"
	c.JSON(http.StatusOK, res)
}

func WriteResponse(res *Response, message string, err error) {
	errMessage := Error{Message: message + "|" + err.Error()}
	res.Error = errMessage
	res.Code = http.StatusBadRequest
	log.Printf("%s %v", message, err)
}
