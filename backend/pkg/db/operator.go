package db

import (
	"SignInHelper/pkg/entity"
	"errors"
	"github.com/go-sql-driver/mysql"
	"log"
)

func (db *client) AddUserInfo(user entity.User) error {
	if err := db.dbClient.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (db *client) InsertNewGithubUser(githubUsername string, token string) error {
	err := db.dbClient.Create(&entity.User{GithubUsername: githubUsername, GithubToken: token, Stared: 1}).Error
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if mysqlErr.Number == 1062 {
			// solve the duplicate key error.
			log.Println("duplicate key, error: " + mysqlErr.Error())
			return errors.New("duplicate key, error: " + mysqlErr.Error())
		} else {
			return errors.New("mysql create error: " + mysqlErr.Error())
		}
	}
	return nil
}

func (db *client) CheckGithubUsernameExist(username string) (bool, error) {
	var count int
	err := db.dbClient.Model(&entity.User{}).Where(&entity.User{GithubUsername: username}).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (db *client) AddSSOInfo(user entity.User) error {
	findUser := db.dbClient.Model(&entity.User{}).Where(&entity.User{GithubUsername: user.GithubUsername})
	err := findUser.Update(&entity.User{
		SSOUsername: user.SSOUsername,
		SSOPassword: user.SSOPassword,
		Phone:       user.Phone,
		Email:       user.Email,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
