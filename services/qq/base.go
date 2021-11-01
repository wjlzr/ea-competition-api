package qq

const (
	accessToken = "oauth2.0/token"
	getOpenId   = "oauth2.0/me"
)

var (
	grantType   = "authorization_code"
	redirectUri = "https%3a%2f%2fwww.wikiglobal.com%2fthird"
)

// 获取token返回
type accessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// 获取token返回
type getOpenIdResponse struct {
	ClientId         string `json:"client_id"`
	OpenId           string `json:"openid"`
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
}
