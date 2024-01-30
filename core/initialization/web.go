package initialization

import (
	"fmt"

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

	if err := common.Web.Run(fmt.Sprintf(":%d", common.Config.System.Port)); err != nil {
		return err
	}

	return
}

// gin router
func initGinRouter() {
	common.Web.GET("/ws", handler.WsHandler)
	common.Web.GET("/health", handler.Health)
}

// 链路追踪
func traceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		c.Request.Header.Set("trace_id", traceID)
		c.Next()
	}
}
