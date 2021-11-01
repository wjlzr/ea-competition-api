package common

import (
	"ea-competition-api/library/i18nresponse"
	"ea-competition-api/services/usercenter"

	"github.com/gin-gonic/gin"
)

// GetCity
func GetCity(c *gin.Context) {

	result, err := usercenter.GetCity(c.ClientIP())
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	if result.Country == "中国" {
		if result.Province == "香港" || result.Province == "澳门" || result.Province == "台湾" {
			result.Country = result.Province
		}
	}
	i18nresponse.Success(c, "ok", result)
}
