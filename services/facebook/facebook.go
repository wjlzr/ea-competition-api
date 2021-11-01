package facebook

import (
	"ea-competition-api/boot/log"
	"ea-competition-api/config"
	"ea-competition-api/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// GetAccessToken 获取token
func GetAccessToken(code string) (*accessTokenResponse, error) {

	request, err := services.Request(http.MethodGet, config.Conf().Fb.Gateway+fmt.Sprintf(accessToken+"?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s", config.Conf().Fb.AppId, config.Conf().Fb.Secret, config.Conf().Fb.RedirectUri, code), nil)
	if err != nil {
		log.Logger().Error("facebook GetAccessToken 请求 err：", zap.Error(err))
		return nil, err
	}

	content, _ := services.ResponseHandle(request)
	var v accessTokenResponse
	_ = json.Unmarshal(content, &v)

	if v.Code != 0 {
		log.Logger().Info("UserCenter GetAccessToken 获取facebook accessToken response：", zap.Any("response", v))
		return nil, errors.New("500")
	}
	return &v, nil
}

// 获取用户信息
func GetMe(accessToken string) (*getMeResponse, error) {

	request, err := services.Request(http.MethodGet, config.Conf().Fb.Gateway+fmt.Sprintf(me+"?access_token=%s", accessToken), nil)
	if err != nil {
		log.Logger().Error("facebook GetMe 请求 err：", zap.Error(err))
		return nil, err
	}

	content, _ := services.ResponseHandle(request)
	var v getMeResponse
	_ = json.Unmarshal(content, &v)

	if v.Error != nil && v.Error.Code != 0 {
		log.Logger().Info("UserCenter GetAccessToken 获取facebook me response：", zap.Any("response", v))
		return nil, errors.New("500")
	}
	return &v, nil
}
