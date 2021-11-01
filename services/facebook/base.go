package facebook

const (
	accessToken = "v10.0/oauth/access_token"
	me          = "me"
)

// 获取token返回
type accessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Message      string `json:"message"`
	Code         int64  `json:"code"`
	ErrorSubCode int64  `json:"error_subcode"`
}

// 获取ID
type getMeResponse struct {
	Name  string              `json:"name"`
	ID    string              `json:"id"`
	Error *getMeErrorResponse `json:"error"`
}

type getMeErrorResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}
