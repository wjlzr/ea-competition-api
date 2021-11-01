package library

import (
	"ea-competition-api/library/constant"
	"ea-competition-api/library/convert/xint64"
	"ea-competition-api/library/token"

	"github.com/gin-gonic/gin"
)

//获取当前平台
//1、web端   2、app端
func GetPlatform(c *gin.Context) uint8 {
	var platform uint8 = 2
	if c.Request.Header.Get(constant.XClient) == constant.Web {
		platform = 1
	}
	return platform
}

//获取token
func GetToken(c *gin.Context) string {
	return c.Request.Header.Get(constant.Authorization)
}

//获取时间戳(秒)
func GetTimetamp(c *gin.Context) int64 {
	return xint64.StrToInt64(c.Request.Header.Get(constant.Timetamp))
}

//获取用户序列号
func GetUserId(c *gin.Context) int64 {
	v, exists := c.Get(constant.UserId)
	if exists {
		return v.(int64)
	}
	return 0
}

//解析用户序列号
func ParseUserID(token *token.Token, c *gin.Context, authorization string) int64 {
	//key := mode.GetRedisKey(constant.RedisAppTokenKey, true)
	//if c.Request.Header.Get(constant.XClient) == constant.Web {
	//	key = mode.GetRedisKey(constant.RedisWebTokenKey, true)
	//}
	if authorization != "" {
		if claims, err := token.Decode(authorization); err == nil {
			//token := db.RedisClusterClient().Get(key + strconv.FormatInt(claims.UserId, 10)).Val()
			//if token == authorization {
			return claims.UserId
			//}
		}
	}
	return 0
}
