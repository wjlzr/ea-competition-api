package awards

import (
	awards2 "ea-competition-api/app/service/awards"
	"ea-competition-api/library/i18nresponse"

	"github.com/gin-gonic/gin"
)

// FindAwards 获取奖励
func FindAwards(c *gin.Context) {

	result, err := awards2.FindAwards()
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", result)
}
