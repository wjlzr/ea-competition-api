package apply

import (
	apply2 "ea-competition-api/app/model/mysql/apply"
	"ea-competition-api/app/model/mysql/oauth"
	"ea-competition-api/app/service/apply"
	"ea-competition-api/library/convert/xint64"
	"ea-competition-api/library/i18nresponse"
	"ea-competition-api/services/skyeyecloud"
	"time"

	"github.com/gin-gonic/gin"
)

// FindApplyInfo 获取报名信息
func FindApplyInfo(c *gin.Context) {

	var req skyeyecloud.StatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := apply.FindApplyInfo(req)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", result)
}

// FindTradeInfo 获取交易商信息
func FindTradeInfo(c *gin.Context) {

	var req oauth.TradeCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := apply.FindTradeInfo(req)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", result)
}

// ApplyAction 报名操作
func ApplyAction(c *gin.Context) {

	var req apply2.Apply
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	success, err := apply.ApplyAction(req)
	if err != nil {
		i18nresponse.Error(c, err.Error())
		return
	}

	i18nresponse.Success(c, "ok", struct {
		Success bool `json:"success"`
	}{Success: success})
}

// FindApplyStatus 获取报名状态
func FindApplyStatus(c *gin.Context) {

	var req apply2.Apply
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	status, err := apply.FindApplyStatus(req)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", struct {
		Status int64 `json:"status"`
	}{Status: status})
}

// UpdateApplyStatus 更新报名参赛资格
func UpdateApplyStatus(c *gin.Context) {

	var req apply2.Apply
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	success, err := apply.UpdateApplyStatus(req)
	if err != nil {
		i18nresponse.Error(c, err.Error())
		return
	}

	i18nresponse.Success(c, "ok", struct {
		Success bool `json:"success"`
	}{Success: success})
}

// FindApply 获取报名信息
func FindApply(c *gin.Context) {

	startTime := c.Request.FormValue("start_time")
	endTime := c.Request.FormValue("end_time")
	qualifications := c.Request.FormValue("qualifications")

	if xint64.StrToInt64(startTime) == 0 {
		i18nresponse.Error(c, "1010004")
		return
	}

	var endTime1 int64

	endTime1 = xint64.StrToInt64(endTime)

	if endTime1 == 0 {
		endTime1 = time.Now().Unix()
	}

	result, err := apply.FindApply(xint64.StrToInt64(startTime), endTime1, xint64.StrToInt64(qualifications))
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", result)
}
