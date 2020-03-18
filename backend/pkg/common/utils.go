package common

import (
	"errors"
	"regexp"
	"strings"
)

func ExtractAccessToken(str string) (string, error) {
	if strings.Contains(str, "error") {
		return "", errors.New("access token format wrong")
	}
	reg := regexp.MustCompile(`=[[:alnum:]]+&`)
	params := reg.FindStringSubmatch(str)
	if len(params) < 1 {
		return "", errors.New("access token format wrong")
	}
	return params[0][1 : len(params[0])-1], nil
}
