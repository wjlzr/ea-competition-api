package usercenter

import "github.com/gin-gonic/gin"

const (
	userName                  = "gsw"
	password                  = "2E6CAA66096F1D4BED0FE21EE5468FE2"
	getToken                  = "api/Permission/Login"                              //获取token
	sendCode                  = "PersonCenter/usercenter/sendcodev3"                //发送验证码
	register                  = "PersonCenter/usercenter/register"                  //用户账号注册
	login                     = "usercenter/loginexpand"                            //账号密码登录
	getUserInfo               = "PersonCenter/usercenter/getuser"                   //获取用户信息
	quickLogin                = "PersonCenter/usercenter/quicklogin"                //快捷登录
	validateCode              = "PersonCenter/usercenter/validatecode"              //验证验证码
	registerValidateUserPhone = "PersonCenter/usercenter/registervalidateuserphone" //验证手机号是否注册过
	thirdPartyLogin           = "PersonCenter/usercenter/thirdpartylogin"           //第三方登录
	validateRegisterPhone     = "PersonCenter/usercenter/validateregisterphone"     //第三方登录 验证手机号是否注册或者第三方是否绑定这个手机号
	thirdPartyRegister        = "PersonCenter/usercenter/thirdpartyregister"        //第三方注册
	geetestRegister           = "geetest/register"                                  //初始化极验
	graphCode                 = "usercenter/getgraphcode"                           //获取验证码
	traderCodes               = "WikiSearch/wikicore/getMultiple"                   //获取交易商信息
	getusermorelist           = "usercenter/getusermorelist"                        //批量返回用户基本信息
	getcity                   = "Third/ip138/getcity"                               //ip反转
)

var (
	authorization   string
	c               *gin.Context
	applicationType = 7
)

type tokenRequest struct {
	UserName string
	Password string
}

// tokenRequest
type tokenResponse struct {
	Status      bool   `json:"status"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// ValidateUserPhoneResponse 验证手机号response new
type ValidateUserPhoneResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Result  string `json:"result"`
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// 发送验证码request
type sendCodeRequest struct {
	AreaCode         string `json:"areaCode"`
	Phone            string `json:"phone"`
	LanguageCode     string `json:"languageCode"`
	UserId           string `json:"userId"`
	SmsBusinessType  int    `json:"smsBusinessType"`
	ApplicationType  int    `json:"applicationType"`
	GeetestValidate  string `json:"geetestValidate"`
	GeetestChallenge string `json:"geetestChallenge"`
	GraphCode        string `json:"graphCode"`
}

// SendCodeResponse 发送验证码response
type SendCodeResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Result struct {
			Requestid string `json:"requestid"`
		} `json:"Result"`
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// CurrencyWithUserResponse 带有用户信息通用返回的参数
type CurrencyWithUserResponse struct {
	Code    int                          `json:"code"`
	Success bool                         `json:"Success"`
	Msg     string                       `json:"msg"`
	Data    CurrencyWithUserResponseData `json:"data"`
}

type CurrencyWithUserResponseData struct {
	Succeed bool                           `json:"succeed"`
	Message string                         `json:"message"`
	Result  CurrencyWithUserResponseResult `json:"result"`
}

type CurrencyWithUserResponseResult struct {
	UserId               string `json:"userId"`
	Nickname             string `json:"nickname"`
	Nick                 string `json:"nick"`
	Avatar               string `json:"avatar"`
	Sex                  int    `json:"sex"`
	Areaflag             string `json:"areaflag"`
	Areacode             string `json:"areacode"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shoppingaddresscount int    `json:"shoppingaddresscount"`
	Realname             string `json:"realname"`
	Isphonecomfirm       bool   `json:"isphonecomfirm"`
	Isemailcomfirm       bool   `json:"isemailcomfirm"`
}

// ValidateCodeResponseOld 验证短信验证码Response old
type ValidateCodeResponseOld struct {
	RequestId string `json:"RequestId"`
	Timestamp string `json:"Timestamp"`
	Content   struct {
		Result struct {
			Succeed bool   `json:"succeed"`
			Message string `json:"message"`
		} `json:"result"`
	} `json:"Content"`
}

// 验证短信验证码Response 用通用CurrencyResponse
//type ValidateCodeResponse struct {
//	Code    int    `json:"code"`
//	Msg     string `json:"msg"`
//	Success bool   `json:"Success"`
//	Data    struct {
//		Succeed bool   `json:"succeed"`
//		Message string `json:"message"`
//	} `json:"Data"`
//}

// CurrencyResponse 通用response new
type CurrencyResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// ValidateRegisterPhoneResponse 验证手机号是否注册或者第三方是否绑定这个手机号response
type ValidateRegisterPhoneResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Result struct {
			ErrorType string `json:"errorType"`
		} `json:"Result"`
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// EvaluationInfoResponse 受评方详情response
type EvaluationInfoResponse struct {
	Code    int                        `json:"code"`
	Msg     string                     `json:"msg"`
	Success bool                       `json:"Success"`
	Data    EvaluationInfoDataResponse `json:"Data"`
}

type EvaluationInfoDataResponse struct {
	Succeed bool        `json:"succeed"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// SpreadCodResponse 交易商点差
type SpreadCodResponse struct {
	Code    int                        `json:"code"`
	Msg     string                     `json:"msg"`
	Success bool                       `json:"Success"`
	Data    EvaluationInfoDataResponse `json:"Data"`
}

// GeetestRegisterActionResponse 初始化极验
type GeetestRegisterActionResponse struct {
	Message string                        `json:"message"`
	Succeed bool                          `json:"succeed"`
	Result  GeetestRegisterResultResponse `json:"result"`
}

// GeetestRegisterResponse 初始化极验
type GeetestRegisterResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Message string                        `json:"message"`
		Succeed bool                          `json:"succeed"`
		Result  GeetestRegisterResultResponse `json:"result"`
	} `json:"Data"`
}

type GeetestRegisterResultResponse struct {
	Challenge  string `json:"challenge"`
	Gt         string `json:"gt"`
	NewCaptcha bool   `json:"new_captcha"`
	Success    int64  `json:"success"`
	Enabled    bool   `json:"enabled"`
}

// GetGraphCodeResponse 获取图形验证码
type GetGraphCodeResponse struct {
	Message string `json:"message"`
	Succeed bool   `json:"succeed"`
	Result  string `json:"result"`
}

// GeetestRegisterResponse 初始化极验
type GetCityResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Succeed bool                      `json:"succeed"`
		Message string                    `json:"message"`
		Result  GetCityDataResultResponse `json:"result"`
	} `json:"Data"`
}

// GetCityResponse ip反转
type GetCityDataResultResponse struct {
	City         string `json:"city"`
	CityCode     string `json:"city_code"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	Ip           string `json:"ip"`
	Operators    string `json:"operators"`
	Province     string `json:"province"`
	ProvinceCode string `json:"province_code"`
}
