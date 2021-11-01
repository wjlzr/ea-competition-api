package user

import (
	"ea-competition-api/app/model/mysql/oauth"
	"ea-competition-api/library"
	"ea-competition-api/library/convert/xint"
	"ea-competition-api/library/convert/xstring"
	"ea-competition-api/library/i18nresponse"
	"ea-competition-api/services/usercenter"

	"github.com/gin-gonic/gin"
)

// GetUserInfo 获取用户详情
func GetUserInfo(c *gin.Context) {

	var req oauth.GetUserInfoRequest
	req.CountryCode = c.Request.FormValue("countryCode")
	req.ApplicationType = xint.StrToInt(c.Request.FormValue("applicationType"))
	if req.CountryCode == "" || req.ApplicationType == 0 {
		i18nresponse.Error(c, "1010004")
		return
	}
	req.UserId = xstring.Int64ToString(library.GetUserId(c))
	result, err := usercenter.GetUserInfo(req)
	if err != nil || result.Succeed != true || result.Message != "success" {
		if result.Message == "" {
			i18nresponse.Error(c, "1010006") // 获取用户信息失败
			return
		} else {
			i18nresponse.Error(c, result.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Result)
}
