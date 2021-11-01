package ucloud

var (
	Action       = "SendUSMSMessage"
	ProjectId    = "org-pk5irq"        // 项目ID
	GnSigContent = "SIG20211026EFDDB1" // 国内短信签名
	//TemplateId   = "UTB210302E20DC2"   // 模板ID
	GgSigContent = "SIG202107068EEBA3" // 国际短信签名
)

type smsResponse struct {
	Action    string `json:"Action"`
	Message   string `json:"Message"`
	RetCode   int64  `json:"RetCode"`
	SessionNo string `json:"SessionNo"`
}
