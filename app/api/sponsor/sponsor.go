package sponsor

import (
	sponsor2 "ea-competition-api/app/service/sponsor"
	"ea-competition-api/library/i18nresponse"

	"github.com/gin-gonic/gin"
)

// FindSponsor 获取赞助商
func FindSponsor(c *gin.Context) {

	result, err := sponsor2.FindSponsor()
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", result)
}
