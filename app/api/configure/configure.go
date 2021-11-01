package configure

import (
	configure2 "ea-competition-api/app/service/configure"
	"ea-competition-api/library/i18nresponse"

	"github.com/gin-gonic/gin"
)

// FindConfigure 获取配置
func FindConfigure(c *gin.Context) {

	result, err := configure2.FindConfigure()
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", result)
}
