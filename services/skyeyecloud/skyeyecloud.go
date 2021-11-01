package skyeyecloud

import (
	"bytes"
	"ea-competition-api/boot/log"
	"ea-competition-api/config"
	"ea-competition-api/services"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

// Status 获取用户大赛资格信息
func Status(r StatusRequest) (*StatusDataResponse, error) {

	jsonStr, _ := json.Marshal(r)
	request, err := services.Request(http.MethodPost, config.Conf().SkyEyeCloud.TestUrl+status, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter Register 请求 err：", zap.Error(err))
		return nil, err
	}

	content, _ := services.ResponseHandle(request)

	var res StatusResponse
	_ = json.Unmarshal(content, &res)
	if res.Code != 200 {
		log.Logger().Info("skyeyecloud Status Error response：", zap.Any("response", res))
		return nil, errors.New("500")
	}
	return &res.Data, nil
}
