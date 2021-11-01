package apply

import (
	"ea-competition-api/app/model/mysql/apply"
	"ea-competition-api/app/model/mysql/applynoticerecord"
	"ea-competition-api/app/model/mysql/oauth"
	"ea-competition-api/services/skyeyecloud"
	"ea-competition-api/services/ucloud"
	"ea-competition-api/services/usercenter"
	"errors"
)

// FindApplyInfo 获取报名相关数据
func FindApplyInfo(request skyeyecloud.StatusRequest) (*skyeyecloud.StatusDataResponse, error) {

	res, err := skyeyecloud.Status(request)
	if err != nil {
		return nil, err
	}

	// 判断是否有参赛资格
	if res.MtInfo == nil {
		return res, nil
	}

	for k, v := range res.MtInfo {
		if v.BrokerType == 1 && v.MtType == 1 && v.AccountType == 1 && v.Currency == "USD" && v.EquityType == 1 {
			res.MtInfo[k].Eligibility = true
		}
	}

	return res, nil
}

// FindTradeInfo 获取交易商信息
func FindTradeInfo(request oauth.TradeCodeRequest) (*skyeyecloud.TradeCodeResponse, error) {

	// codes, mapSort := xjson.JsonToString(config.Conf().SkyEyeCloud.TradeJsonUrl, "brokerid")
	request.Codes = []string{
		"0001264568",
		"1670082012",
		"8421818926",
		"0001461138",
		"0001390005",
		"5991606397",
		"5121914749",
		"3351410785",
		"6251173673",
	}

	res, err := usercenter.GeTradeInfo(request)
	if err != nil {
		return nil, err
	}
	var mapBrokers map[string]int = make(map[string]int)
	for k, v := range request.Codes {
		mapBrokers[v] = k + 1
	}
	// 排序
	for k, v := range res.Data.Result {
		// var sort int
		// for k1, v1 := range request.Codes {
		// 	if v.Code == v1 {
		// 		sort = k1 + 1
		// 	}
		// }
		res.Data.Result[k].Sort = mapBrokers[v.Code] //sort
	}

	return res, nil
}

// ApplyAction 报名
func ApplyAction(apply apply.Apply) (bool, error) {

	// 判断是否已报名
	_, count, err := apply.FindOne()
	if err != nil {
		return false, errors.New("500")
	}
	if count != 0 {
		return false, errors.New("1010016")
	}

	// 入库
	_, err = apply.Create()
	if err != nil {
		return false, errors.New("500")
	}

	// 查询用户信息
	var req oauth.UserListRequest
	req.UserIds = []string{apply.UserId}
	res, err := usercenter.GetUserList(req)
	if err != nil || len(res.Result) == 0 {
		return true, nil
	}

	// 发送短信通知 todo 后期根据业务量可改为消息队列异步发送
	result, err := ucloud.SmsSend("("+res.Result[0].AreaCode+")", res.Result[0].PhoneNumber)
	if err != nil && result != nil {
		if err.Error() == "100000" {
			var applynoticerecord applynoticerecord.ApplyNoticeRecord
			applynoticerecord.AreaCode = res.Result[0].AreaCode
			applynoticerecord.Communication = res.Result[0].PhoneNumber
			applynoticerecord.UserId = res.Result[0].UserId
			applynoticerecord.Type = 1
			applynoticerecord.Result = 2
			applynoticerecord.Message = result.Message
			applynoticerecord.SessionNo = result.SessionNo
			applynoticerecord.RetCode = result.RetCode
			_, _ = applynoticerecord.Create()
		}
	}

	return true, nil
}

// FindApplyStatus 查询报名状态
func FindApplyStatus(apply apply.Apply) (int64, error) {

	// 查询状态
	res, count, err := apply.FindOne()
	if err != nil {
		return 0, err
	}

	// 未报名
	if count == 0 {
		return 0, nil
	}
	// 0 未报名 1 已报名 不符合参赛资格 2 已报名 符合参赛资格
	if res.Qualifications == 2 {
		return 1, nil
	}

	return 2, nil
}

// UpdateApplyStatus 查询报名状态
func UpdateApplyStatus(apply1 apply.Apply) (bool, error) {

	var request skyeyecloud.StatusRequest
	request.UserId = apply1.UserId

	// 先判断此人是否有参赛资格
	res, err := skyeyecloud.Status(request)
	if err != nil {
		return false, errors.New("500")
	}

	// 判断是否有参赛资格
	if res.Vps == 0 || res.Available1 <= 0 || res.Available3 == 1 {
		return false, errors.New("1010015")
	}

	// 查询状态
	_, err = apply1.Update(map[string]interface{}{"user_id": request.UserId})
	if err != nil {
		return false, errors.New("500")
	}

	return true, nil
}

// FindApply 查询报名数据
func FindApply(startTime, endTime, qualifications int64) ([]apply.Apply, error) {

	result, err := apply.Apply{}.QueryStatistics(startTime, endTime, qualifications)
	if err != nil {
		return nil, err
	}
	return result, nil
}
