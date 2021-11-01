package ucloud

import (
	"ea-competition-api/boot/log"
	"ea-competition-api/config"
	"ea-competition-api/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"go.uber.org/zap"
)

func SmsSend(areaCode, phone string) (*smsResponse, error) {

	sigContent := getSigContent(areaCode)

	template := getTemplate(areaCode)

	q := "Action=" + Action + "&ProjectId=" + ProjectId + "&PhoneNumbers.0=" + areaCode + phone + "&SigContent=" + sigContent + "&TemplateId=" + template + "&PublicKey=" + config.Conf().UCloud.PublicKey
	sign := Sign(q)

	request, err := services.Request(http.MethodGet, config.Conf().UCloud.Gateway+fmt.Sprintf("?Action=%s&ProjectId=%s&PhoneNumbers.0=%s&SigContent=%s&TemplateId=%s&PublicKey=%s&Signature=%s", Action, ProjectId, areaCode+phone, sigContent, template, config.Conf().UCloud.PublicKey, sign), nil)
	if err != nil {
		log.Logger().Error("ucloud SmsSend 请求 err：", zap.Error(err))
		return nil, err
	}
	content, _ := services.ResponseHandle(request)
	var v smsResponse
	_ = json.Unmarshal(content, &v)
	if v.RetCode != 0 {
		log.Logger().Info("ucloud SmsSend 发送短信 response：", zap.Any("response", v))
		return &v, errors.New("100000")
	}
	return &v, nil
}

func Sign(query string) string {

	cred := &auth.Credential{
		PublicKey:  config.Conf().UCloud.PublicKey,
		PrivateKey: config.Conf().UCloud.PrivateKey,
	}
	//pp.Println(cred.CreateSign(query))
	return cred.CreateSign(query)
}

func getTemplate(areaCode string) string {

	// 默认英语
	template := "UTN21102629E6A6"
	switch areaCode {
	case "(0033)":
		// 法语
		template = "UTB210305CAF29B"
	case "(0034)":
		// 西班牙
		template = "UTB2103059C8411"
	case "(0039)":
		template = "UTA210305E74E23"
	case "(0049)":
		template = "UTA2103051C8DDA"
	case "(0062)":
		template = "UTB2103057C8E5A"
	case "(0063)":
		template = "UTA2103051737BB"
	case "(0066)":
		template = "UTB210305A019E1"
	case "(007)":
		template = "UTB210305C9FD5F"
	case "(0081)":
		template = "UTB21030528A1C3"
	case "(0084)":
		template = "UTB2103053BDEAC"
	default:
		template = "UTN21102629E6A6"
	}
	return template
}

func getSigContent(areaCode string) string {

	if areaCode == "0086" {
		return GnSigContent
	}

	return GgSigContent
}
