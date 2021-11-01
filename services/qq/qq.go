package qq

import (
	"ea-competition-api/boot/log"
	"ea-competition-api/config"
	"ea-competition-api/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/k0kubun/pp"
	"go.uber.org/zap"
)

func GetAccessToken(code string) (*accessTokenResponse, error) {

	request, err := services.Request(http.MethodGet, config.Conf().Qq.Gateway+fmt.Sprintf(accessToken+"?client_id=%s&client_secret=%s&code=%s&grant_type=%s&redirect_uri=%s", config.Conf().Qq.AppId, config.Conf().Qq.AppKey, code, grantType, redirectUri), nil)
	if err != nil {
		log.Logger().Error("qq GetAccessToken 请求 err：", zap.Error(err))
		return nil, err
	}

	content, _ := services.ResponseHandle(request)
	str := string(content)
	var v accessTokenResponse
	if str[0:1] == "c" {
		count := utf8.RuneCountInString(str)
		str1 := str[10 : count-3]
		_ = json.Unmarshal([]byte(str1), &v)
	} else if str[0:1] == "a" {
		str1 := str[13:45]
		v.AccessToken = str1
	}
	if v.Error != 0 || v.ErrorDescription != "" {
		log.Logger().Info("UserCenter GetAccessToken 获取QQ accessToken response：", zap.Any("response", v))
		return nil, errors.New("500")
	}
	return &v, nil
}

func GetOpenId(accessToken string) (*getOpenIdResponse, error) {

	_, _ = pp.Println("access_token")
	_, _ = pp.Println(accessToken)
	request, err := services.Request(http.MethodGet, config.Conf().Qq.Gateway+fmt.Sprintf(getOpenId+"?access_token=%s", accessToken), nil)
	if err != nil {
		log.Logger().Error("qq GetOpenId 请求 err：", zap.Error(err))
		return nil, err
	}

	content, _ := services.ResponseHandle(request)
	pp.Println(string(content))
	var v getOpenIdResponse
	count := utf8.RuneCountInString(string(content))
	str := string(content)[10 : count-3]
	_ = json.Unmarshal([]byte(str), &v)
	pp.Println(v)
	if v.Error != 0 || v.ErrorDescription != "" {
		log.Logger().Info("UserCenter GetOpenId 获取QQ openid response：", zap.Any("response", v))
		return nil, errors.New("500")
	}
	return &v, nil
}
