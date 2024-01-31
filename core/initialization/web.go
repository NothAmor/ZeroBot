package initialization

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 初始化Gin Web
func initGin() (err error) {
	if common.Config.System.Environment == common.SysEnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	common.Web = gin.Default()

	common.Web.Use(traceMiddleware())
	initGinRouter()

	return
}

// gin router
func initGinRouter() {
	common.Web.GET("/", handler.WsHandler)
	common.Web.GET("/health", handler.Health)
}

// 链路追踪
func traceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		c.Set("traceID", traceID)
		c.Next()
	}
}
