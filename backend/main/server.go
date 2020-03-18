package main

import (
	"SignInHelper/pkg/service"
	"SignInHelper/pkg/service/recordusers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	g := r.Group("/hitcarder")
	g.Any("/ping", service.Ping)
	{
		h := recordusers.New()
		g.GET("/record", h.RecordUser)
		g.GET("/auth_check", h.AuthCheck)
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
