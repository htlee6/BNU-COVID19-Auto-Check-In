package db

import (
	"SignInHelper/pkg/entity"
	"SignInHelper/pkg/secrets"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Client interface {
	AddUserInfo(user entity.User) error
	InsertNewGithubUser(githubUsername string, token string) error
	CheckGithubUsernameExist(username string) (bool, error)
	AddSSOInfo(user entity.User) error
}

type client struct {
	dbClient *gorm.DB
}

func NewClient() Client {
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		secrets.MYSQL_USERNAME, secrets.MYSQL_PASSQWORD, secrets.MYSQL_HOST, secrets.MYSQL_PORT, secrets.MYSQL_DATABASE_NCOV,
	)

	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("sql open failed []", err)
	}
	db.SingularTable(true)
	return &client{dbClient: db}
}
