package oauth

import (
	"ea-competition-api/app/model/mysql/oauth"
	oauth1 "ea-competition-api/app/model/mysql/oauth"
	user2 "ea-competition-api/app/model/mysql/user"
	"ea-competition-api/library/convert/xint"
	"ea-competition-api/library/convert/xint64"
	"ea-competition-api/library/i18nresponse"
	"ea-competition-api/library/token"
	"ea-competition-api/services/facebook"
	"encoding/json"
	"time"

	"ea-competition-api/services/qq"
	"ea-competition-api/services/usercenter"
	"ea-competition-api/services/wechat"

	"github.com/k0kubun/pp"

	"github.com/gin-gonic/gin"
)

//发送验证码
func SmsSend(c *gin.Context) {

	code := c.Request.FormValue("code")
	phone := c.Request.FormValue("phone")
	languageCode := c.Request.FormValue("languageCode")
	userId := c.Request.FormValue("userId")
	smsBusinessType := xint.StrToInt(c.Request.FormValue("smsBusinessType"))
	geetestValidate := c.Request.FormValue("GeetestValidate")
	geetestChallenge := c.Request.FormValue("GeetestChallenge")
	graphCode := c.Request.FormValue("graphCode")

	result, err := usercenter.ValidatePhone(code, phone)
	if err != nil {
		i18nresponse.Error(c, "1010001") //手机号验证失败
		return
	}

	// 行为判断 1-登录,2-注册,3-第三方,4-验证手机号,5-修改手机号,6-忘 记密码,7-实名认证
	if smsBusinessType == 0 {
		smsBusinessType = 2
		if result == false {
			smsBusinessType = 1
		}
	}

	// 发送验证码
	smsResult, err := usercenter.SendCode(code, phone, languageCode, userId, geetestValidate, geetestChallenge, graphCode, smsBusinessType)
	if err != nil || smsResult.Data.Succeed != true {
		i18nresponse.Error(c, "1010002") // 发送验证码失败,请重试
		return
	}

	i18nresponse.Success(c, "ok", struct {
		Success   bool   `json:"success"`
		Requestid string `json:"requestid"`
	}{Success: result, Requestid: smsResult.Data.Result.Requestid})
}

//账号密码登录
func Login(c *gin.Context) {

	var t oauth1.TokenAndUserInfoResponse
	var req oauth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := usercenter.Login(req)
	if err != nil || result.Succeed != true {
		if result.Message == "" {
			i18nresponse.Error(c, "1010003") //登录失败,请重试
			return
		} else {
			i18nresponse.Error(c, result.Message)
			return
		}
	}

	if result.Result.UserId != "" {
		t.IsRegister = true
	}

	jsonStu, err := json.Marshal(result.Result)
	_ = json.Unmarshal(jsonStu, &t.UserInfo)
	t.Authorization = getToken(result.Result.UserId)
	//t.Authorization, err = t2.Encode(convert.StrToInt64(result.Result.UserId), time.Now().Add(time.Second*180).Unix())
	i18nresponse.Success(c, "ok", t)
}

//快捷登录
func QuickLogin(c *gin.Context) {

	var t oauth.TokenAndUserInfoResponse
	var req oauth.QuickLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := usercenter.QuickLogin(req)
	if err != nil || result.Succeed != true {
		if result.Message == "" {
			i18nresponse.Error(c, "1010003") //登录失败,请重试
			return
		} else {
			i18nresponse.Error(c, result.Message)
			return
		}
	}

	jsonStu, err := json.Marshal(result.Result)
	_ = json.Unmarshal(jsonStu, &t.UserInfo)
	t.Authorization = getToken(result.Result.UserId)
	//t.Authorization, err = t2.Encode(convert.StrToInt64(result.Result.UserId), time.Now().Add(time.Second*180).Unix())

	i18nresponse.Success(c, "ok", t)
}

//注册
func Register(c *gin.Context) {

	var req oauth.RegisterRequest
	var t oauth.TokenAndUserInfoResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := usercenter.Register(req)
	if err != nil || result.Succeed != true {
		if result.Message == "" {
			i18nresponse.Error(c, "_S000018") // 注册失败，请联系客服
			return
		} else {
			i18nresponse.Error(c, result.Message)
			return
		}
	}

	// 注册成功后记录用户关键信息
	var user user2.User
	user.OpenId = ""
	user.UserId = result.Result.UserId
	user.CountryCode = req.CountryCode
	user.AreaCode = req.AreaCode
	user.Phone = req.Phone
	user.Password = req.Password
	user.Email = result.Result.Email
	user.AccountType = "1000" // 1000 为普通注册
	user.Password = result.Result.Password
	_, err = user.Create()
	if err != nil {
		i18nresponse.Error(c, "_S000018") // 注册失败，请联系客服
		return
	}

	jsonStu, err := json.Marshal(result.Result)
	_ = json.Unmarshal(jsonStu, &t.UserInfo)
	t.Authorization = getToken(result.Result.UserId)
	//t.Authorization, err = t2.Encode(convert.StrToInt64(result.Result.UserId), time.Now().Add(time.Second*180).Unix())

	i18nresponse.Success(c, "ok", t)
}

// 验证短信验证码
func ValidateCode(c *gin.Context) {

	var req oauth.ValidateCodeRequest
	req.AreaCode = c.Request.FormValue("areaCode")
	req.Smscode = c.Request.FormValue("smscode")
	req.PhoneNumber = c.Request.FormValue("phoneNumber")
	if req.AreaCode == "" || req.Smscode == "" || req.PhoneNumber == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := usercenter.ValidateCode(req)
	if err != nil {
		i18nresponse.Error(c, "_S000029")
		return
	}

	if result.Success == false || result.Data.Succeed == false {
		if result.Data.Message == "" {
			i18nresponse.Error(c, "1010005") // 验证码校验失败,请重试
			return
		} else {
			i18nresponse.Error(c, result.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Data)
}

// GetOpenId 获取微信openid
func GetWeChatOpenId(c *gin.Context) {

	code := c.Request.FormValue("code")
	if code == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	// 获取token
	r, err := wechat.GetAccessToken(code)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", struct {
		OpenId string `json:"open_id"`
	}{OpenId: r.UnionId})
}

// 微信登录
func ThirdPartyLogin(c *gin.Context) {

	var req oauth.ThirdPartyLoginRequest
	var t oauth.TokenAndUserInfoResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	req.Ip = c.ClientIP()
	result, err := usercenter.ThirdPartyLogin(req)
	if err != nil {
		i18nresponse.Error(c, "1010003")
		return
	}

	if result.Succeed != true {
		i18nresponse.Error(c, "1010003")
		return
	}

	if result.Succeed == true && result.Result.UserId != "" {
		// 已注册可以正常登录
		jsonStu, _ := json.Marshal(result.Result)
		_ = json.Unmarshal(jsonStu, &t.UserInfo)
		t.Authorization = getToken(result.Result.UserId)
		//t.Authorization, err = t2.Encode(convert.StrToInt64(result.Result.UserId), time.Now().Add(time.Second*180).Unix())
		i18nresponse.Success(c, "ok", t)
		return
	}

	// 未注册走注册流程
	i18nresponse.Success(c, "ok", nil)
}

// 单纯的发送验证码
func SendCode(c *gin.Context) {

	code := c.Request.FormValue("code")
	phone := c.Request.FormValue("phone")
	languageCode := c.Request.FormValue("languageCode")
	smsBusinessType := xint.StrToInt(c.Request.FormValue("smsBusinessType"))
	geetestValidate := c.Request.FormValue("GeetestValidate")
	geetestChallenge := c.Request.FormValue("GeetestChallenge")
	graphCode := c.Request.FormValue("graphCode")

	// 发送验证码
	smsResult, err := usercenter.SendCode(code, phone, languageCode, "", geetestValidate, geetestChallenge, graphCode, smsBusinessType)
	if err != nil || smsResult.Data.Succeed != true {
		if smsResult.Data.Message == "" {
			i18nresponse.Error(c, "1010002") // 发送验证码失败,请重试
			return
		} else {
			i18nresponse.Error(c, smsResult.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", struct {
		Success   bool   `json:"success"`
		Requestid string `json:"requestid"`
	}{Success: true, Requestid: smsResult.Data.Result.Requestid})
}

// ValidateRegisterPhone  第三方登录-验证手机号是否注册或者第三方是否绑定这个手机号
func ValidateRegisterPhone(c *gin.Context) {

	var req oauth.ValidateRegisterPhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := usercenter.ValidateRegisterPhone(req)
	if err != nil || result.Data.Succeed == false {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", struct {
		Type string `json:"type"`
	}{Type: result.Data.Result.ErrorType})
}

// 第三方注册
func ThirdRegister(c *gin.Context) {

	var req oauth.ThirdRegisterRequest
	var t oauth.TokenAndUserInfoResponse
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := usercenter.ThirdRegister(req)
	if err != nil || result.Succeed != true {
		if result.Message == "" {
			i18nresponse.Error(c, "_S000018") // 注册失败，请联系客服
			return
		} else {
			i18nresponse.Error(c, result.Message)
			return
		}
	}

	// 注册成功后记录用户关键信息
	var user user2.User
	user.OpenId = req.OpenId
	user.UserId = result.Result.UserId
	user.CountryCode = req.CountryCode
	user.AreaCode = req.AreaCode
	user.Phone = req.Phone
	user.Password = req.Password
	user.Email = result.Result.Email
	user.AccountType = req.AccountType
	user.Password = result.Result.Password
	_, err = user.Create()
	if err != nil {
		i18nresponse.Error(c, "_S000018") // 注册失败，请联系客服
		return
	}

	jsonStu, err := json.Marshal(result.Result)
	_ = json.Unmarshal(jsonStu, &t.UserInfo)
	t.Authorization = getToken(result.Result.UserId)
	//t.Authorization, err = t2.Encode(convert.StrToInt64(result.Result.UserId), time.Now().Add(time.Second*180).Unix())
	i18nresponse.Success(c, "ok", t)
}

// GetQqOpenId 获取QQ openid
func GetQqOpenId(c *gin.Context) {

	code := c.Request.FormValue("code")
	if code == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	// 获取token
	r, err := qq.GetAccessToken(code)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	// 获取openid
	r2, err := qq.GetOpenId(r.AccessToken)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	pp.Println("QQ 获取OpenId")
	pp.Println(r2)
	i18nresponse.Success(c, "ok", struct {
		OpenId string `json:"open_id"`
	}{OpenId: r2.OpenId})
}

// GetFbOpenId 获取facebook openid
func GetFbOpenId(c *gin.Context) {

	code := c.Request.FormValue("code")
	if code == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	// 获取token
	r, err := facebook.GetAccessToken(code)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	// 获取openid
	r2, err := facebook.GetMe(r.AccessToken)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", struct {
		OpenId string `json:"open_id"`
	}{OpenId: r2.ID})
}

// Geetest 初始化极验
func Geetest(c *gin.Context) {

	clientType := c.Request.FormValue("clienttype")
	if clientType == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	// 获取openid
	challenge, err := usercenter.Geetest(clientType)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", challenge)
}

// GetGraphCode 获取图片验证码
func GetGraphCode(c *gin.Context) {

	var req oauth.GraphCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := usercenter.GetGraphCode(req)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", struct {
		Result string `json:"result"`
	}{Result: result})
}

//账号密码登录
func LoginV2(c *gin.Context) {

	var req oauth.LoginV2Request
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	authorization := getToken(req.UserId)
	i18nresponse.Success(c, "ok", struct {
		Authorization string `json:"authorization"`
	}{Authorization: authorization})
}

func getToken(userId string) string {

	t2 := token.Token{}
	tokenStr, _ := t2.Encode(xint64.StrToInt64(userId), time.Now().Add(time.Second*600).Unix())
	return tokenStr
}
