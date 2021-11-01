package oauth

// 注册请求参数
type RegisterRequest struct {
	AreaFlag             string `json:"areaFlag"`
	AreaCode             string `json:"areaCode"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Sex                  int    `json:"sex"`
	Lastname             string `json:"lastname"`
	IsSkip               int    `json:"isSkip"`
	ApplicationType      int    `json:"applicationType"`
	LanguageCode         string `json:"languageCode"`
	CountryCode          string `json:"countryCode"`
	Version              string `json:"version"`
	Ip                   string `json:"ip"`
	RequestId            string `json:"requestId"`
	RegistrationPlatform int    `json:"registrationPlatform"`
	DeviceInformation    string `json:"deviceInformation"`
	DeviceCode           string `json:"deviceCode"`
	UserFirstName        string `json:"userFirstName"`
}

// 注册请求参数
type ThirdRegisterRequest struct {
	OpenId               string `json:"openId"`
	AccountType          string `json:"accountType"`
	AccountNick          string `json:"accountNick"`
	AccountHead          string `json:"accountHead"`
	Sex                  int    `json:"sex"`
	Lastname             string `json:"lastname"`
	Phone                string `json:"phone"`
	AreaFlag             string `json:"areaFlag"`
	AreaCode             string `json:"areaCode"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	ApplicationType      int    `json:"applicationType"`
	CountryCode          string `json:"countryCode"`
	LanguageCode         string `json:"languageCode"`
	Version              string `json:"version"`
	Ip                   string `json:"ip"`
	IsSkip               int    `json:"isSkip"`
	RequestId            string `json:"requestId"`
	RegistrationPlatform int    `json:"registrationPlatform"`
	DeviceInformation    string `json:"deviceInformation"`
	UserFirstName        string `json:"userFirstName"`
	MsgCode              string `json:"msgCode"`
}

// 短信验证码验证
type ValidateCodeRequest struct {
	AreaCode    string `json:"areaCode"`
	PhoneNumber string `json:"phoneNumber"`
	Smscode     string `json:"smscode"`
	UserId      string `json:"userId"`
}

// 快捷登录Request
type QuickLoginRequest struct {
	AreaCode        string `json:"areaCode"`
	LanguageCode    string `json:"languageCode"`
	Phone           string `json:"phone"`
	MsgCode         string `json:"msgCode"`
	ApplicationType int    `json:"applicationType"`
	Ip              string `json:"ip"`
	Version         string `json:"version"`
	EquipmentType   int    `json:"equipmentType"`
}

// 账号密码登录Request
type LoginRequest struct {
	Account         string `json:"account"`
	Password        string `json:"password"`
	LanguageCode    string `json:"languageCode"`
	CountryCode     string `json:"countryCode"`
	Ip              string `json:"ip"`
	EquipmentType   int    `json:"equipmentType"`
	ApplicationType int    `json:"applicationType"`
}

// 通过手机号找回密码Request
type ModifyPassByPhoneRequest struct {
	UserId    string `json:"userId"`
	AreaCode  string `json:"areaCode"`
	Phone     string `json:"phone"`
	Npwd      string `json:"npwd"`
	RequestId string `json:"requestId"`
}

// 通过旧密码改新密码Request
type ModifyPassByOldRequest struct {
	UserId string `json:"userId"`
	Opwd   string `json:"opwd"`
	Npwd   string `json:"npwd"`
}

// 发送邮箱验证码Request
type SendEmailCodeRequest struct {
	Email           string `json:"email"`
	UserId          string `json:"userId"`
	EmailType       int    `json:"emailType"`
	LanguageCode    string `json:"languageCode"`
	ApplicationType int    `json:"applicationType"`
}

// 验证邮箱（验证码）
type ConfirmEmailByCodeRequest struct {
	Email           string `json:"email"`
	UserId          string `json:"userId"`
	Code            string `json:"code"`
	ApplicationType int    `json:"applicationType"`
}

// 验证邮箱（链接）
type ConfirmEmailByLineRequest struct {
	Email           string `json:"email"`
	UserId          string `json:"userId"`
	ApplicationType int    `json:"applicationType"`
}

// 第三方登录
type ThirdPartyLoginRequest struct {
	Id              string `json:"id"`          // 第三方openid
	AccountType     string `json:"accountType"` // 351：QQ,352：微信,361:推特,371:Facebook,381:GMail
	ApplicationType int    `json:"applicationType"`
	Ip              string `json:"ip"` // 注册Ip
}

//  第三方登录-验证手机号是否注册或者第三方是否绑定这个手机号
type ValidateRegisterPhoneRequest struct {
	AreaCode        string `json:"areaCode"`
	Phone           string `json:"phone"`
	UnionId         string `json:"unionId"`     // 第三方openid
	AccountType     string `json:"accountType"` // 351：QQ,352：微信,361:推特,371:Facebook,381:GMail
	ApplicationType int    `json:"applicationType"`
}

// 获取微信CODE
type WeChatCodeRequest struct {
	RedirectUri  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"` // code
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

// 校验邮箱是否验证Request
type CheckMailboxRequest struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

//获取用户信息Request
type GetUserInfoRequest struct {
	UserId          string `json:"userId"`
	CountryCode     string `json:"countryCode"`
	ApplicationType int    `json:"applicationType"`
}

// token
type TokenResponse struct {
	Authorization string `json:"authorization"`
}

// tokenAndUserInfo
type TokenAndUserInfoResponse struct {
	UserInfo      userInfo `json:"user_info"`
	Authorization string   `json:"authorization"`
	IsRegister    bool     `json:"is_register"`
}

type userInfo struct {
	UserId               string `json:"userId"`
	Nickname             string `json:"nickname"`
	Avatar               string `json:"avatar"`
	Sex                  int    `json:"sex"`
	Areaflag             string `json:"areaflag"`
	Areacode             string `json:"areacode"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	Shoppingaddresscount int    `json:"shoppingaddresscount"`
	Realname             string `json:"realname"`
	Isphonecomfirm       bool   `json:"isphonecomfirm"`
	Isemailcomfirm       bool   `json:"isemailcomfirm"`
}

type Apple struct {
	Code   string `json:"code"`
	UserId string `json:"user_id"`
}

// 获取图形验证码request
type GraphCodeRequest struct {
	AreaCode    string `json:"areaCode"`
	PhoneNumber string `json:"phonenumber"`
	SmsSignName int64  `json:"smsSignName"`
	Length      int64  `json:"length"`
	Width       int64  `json:"width"`
	Heigth      int64  `json:"heigth"`
}

// 获取交易商信息
type TradeCodeRequest struct {
	LanguageCode string   `json:"language"`
	CountryCode  string   `json:"country"`
	Codes        []string `json:"codes"`
}

// 获取交易商信息
type UserListRequest struct {
	UserIds []string `json:"userIds"`
}

type LoginV2Request struct {
	UserId string `json:"user_id" binding:"required"`
}
