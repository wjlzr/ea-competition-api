package router

import (
	"ea-competition-api/app/api/apply"
	"ea-competition-api/app/api/awards"
	"ea-competition-api/app/api/common"
	"ea-competition-api/app/api/configure"
	"ea-competition-api/app/api/oauth"
	"ea-competition-api/app/api/sponsor"
	"ea-competition-api/app/api/user"
	"ea-competition-api/config"
	"ea-competition-api/library/token"
	"ea-competition-api/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 路由配置
func InitRouter(token *token.Token, zapLogger *zap.Logger) *gin.Engine {

	engine := gin.New()
	//gin中间件加载
	engine.Use(middleware.Cors())
	engine.Use(middleware.Secure())
	engine.Use(middleware.Language())
	engine.Use(middleware.Ginzap(zapLogger, time.RFC3339, true))
	engine.Use(middleware.RecoveryWithZap(zapLogger, true))

	engine.Use(middleware.UserAuthMiddleware(token, middleware.AllowPathPrefixSkipper(config.Conf().API.AllowPathPrefixSkipper)))

	apiGroup := engine.Group("api/v1")

	//oauth
	oauthGroup(apiGroup.Group("oauth"))

	// 获取配置
	apiGroup.GET("/findConfigure", configure.FindConfigure)
	// 获取奖励
	apiGroup.GET("/findAwards", awards.FindAwards)
	// 获取赞助商
	apiGroup.GET("/findSponsor", sponsor.FindSponsor)
	// 获取用户信息
	apiGroup.GET("/user/info", user.GetUserInfo)
	// 获取报名信息
	apiGroup.POST("/findApplyInfo", apply.FindApplyInfo)
	// 获取交易商信息
	apiGroup.POST("/findTradeInfo", apply.FindTradeInfo)
	// 报名
	apiGroup.POST("/apply", apply.ApplyAction)
	// 获取报名状态
	apiGroup.POST("/apply/FindApplyStatus", apply.FindApplyStatus)
	// 更改参赛资格
	apiGroup.POST("/apply/UpdateApplyStatus", apply.UpdateApplyStatus)
	// 获取统计数据
	apiGroup.GET("/findApply", apply.FindApply)
	// 获取统计数据
	apiGroup.GET("/getCity", common.GetCity)

	return engine
}

// 认证路由组
func oauthGroup(rg *gin.RouterGroup) {
	//登录
	rg.POST("/login", oauth.Login)
	rg.POST("/loginV2", oauth.LoginV2)
	rg.GET("/smsSend", oauth.SmsSend)
	rg.POST("/register", oauth.Register)
	rg.GET("/validateCode", oauth.ValidateCode)
	rg.POST("/quickLogin", oauth.QuickLogin)
	// 微信登录
	rg.POST("/thirdPartyLogin", oauth.ThirdPartyLogin)
	rg.GET("/sendCode", oauth.SendCode)
	rg.GET("/getWeChatOpenId", oauth.GetWeChatOpenId)
	rg.POST("/validateRegisterPhone", oauth.ValidateRegisterPhone)
	rg.POST("/thirdRegister", oauth.ThirdRegister)
	rg.GET("/getQqOpenId", oauth.GetQqOpenId)
	rg.GET("/getFbOpenId", oauth.GetFbOpenId)
	rg.GET("/geetest", oauth.Geetest)
	rg.POST("/getGraphCode", oauth.GetGraphCode)
}
