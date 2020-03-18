package rpc

import (
	"SignInHelper/pkg/common"
	"SignInHelper/pkg/secrets"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAccessToken(code string) (string, error) {
	request := getAccessTokenRequest{
		ClientID:     secrets.CLIENT_ID,
		ClientSecret: secrets.CLIENT_SECRET,
		Code:         code,
	}
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	resp, err := http.Post("https://github.com/login/oauth/access_token", "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		return "", err
	}

	fmt.Println("resp:", resp)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("body: ", string(body))
	accessToken, err := common.ExtractAccessToken(string(body))
	if err != nil {
		return "", err
	}
	fmt.Println("token: ", accessToken)
	return accessToken, nil
}

func GetUserName(accessToken string) (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	request.Header.Set("Authorization", "token "+accessToken)
	if err != nil {
		return "", err
	}
	resp, _ := client.Do(request)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	response := &getUserInfoResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return "", err
	}
	return response.Username, nil
}

func StarProject(accessToken string) error {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", "https://api.github.com/user/starred/timfaner/BUAA-nCov-Hitcarder", nil)
	request.Header.Set("Content-Length", "0")
	request.Header.Set("Authorization", "token "+accessToken)
	if err != nil {
		return err
	}
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	return nil
}
