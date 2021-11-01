package skyeyecloud

var status = "competition/sign/status"

// StatusRequest
type StatusRequest struct {
	UserId   string `json:"user_id"`
	Code     string `json:"code"`
	PageNum  int64  `json:"page_num"`
	PageSize int64  `json:"page_size"`
}

// 状态返回参数
type StatusResponse struct {
	Code int64              `json:"code"`
	Data StatusDataResponse `json:"data"`
}

type StatusDataResponse struct {
	Vps        int64                  `json:"vps"`
	Phone      string                 `json:"phone"`
	Min        int64                  `json:"min"`
	Total      int64                  `json:"total"`
	Available1 int64                  `json:"available1"`
	Available2 int64                  `json:"available2"`
	Available3 int64                  `json:"available3"`
	MtInfo     []statusMtInfoResponse `json:"mtinfo"`
}

type statusMtInfoResponse struct {
	MtId         int64   `json:"mtid"`
	MtAccount    int64   `json:"mt_account"`
	BrokerName   string  `json:"broker_name"`
	BrokerIcon   string  `json:"broker_icon"`
	BrokerType   int64   `json:"broker_type"`
	MtType       int64   `json:"mt_type"`
	AccountType  int64   `json:"account_type"`
	Currency     string  `json:"currency"`
	CurrencyType int64   `json:"currency_type"`
	Equity       float64 `json:"equity"`
	EquityType   int64   `json:"equity_type"`
	Time         int64   `json:"time"`
	Eligibility  bool    `json:"eligibility"`
}

// 获取交易商返回值
type TradeCodeResponse struct {
	Code int64                 `json:"code"`
	Msg  string                `json:"msg"`
	Data TradeCodeDataResponse `json:"Data"`
}

type TradeCodeDataResponse struct {
	Succeed bool                          `json:"succeed"`
	Message string                        `json:"message"`
	Result  []TradeCodeDataResultResponse `json:"result"`
}

type TradeCodeDataResultResponse struct {
	Code           string `json:"code"`
	Ico            string `json:"ico"`
	LocalShortName string `json:"localShortName"`
	Sort           int    `json:"sort"`
}

type UserListResponse struct {
	Code int64                `json:"code"`
	Msg  string               `json:"msg"`
	Data UserListDataResponse `json:"Data"`
}

type UserListDataResponse struct {
	Succeed bool                         `json:"succeed"`
	Message string                       `json:"message"`
	Result  []UserListDataResultResponse `json:"result"`
}

type UserListDataResultResponse struct {
	UserId      string `json:"userId"`
	AreaCode    string `json:"areaCode"`
	PhoneNumber string `json:"phoneNumber"`
}
