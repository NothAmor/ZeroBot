package initialization

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/sirupsen/logrus"
)

// 初始化logger
func initLog() (err error) {
	common.Log = logrus.New()

	// 设置日志输出
	stdOutput := os.Stdout
	logFileName := fmt.Sprintf("logs/zeroBot-%s.log", time.Now().Format("2006-01-02"))
	fileOutput, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	common.Log.SetOutput(io.MultiWriter(stdOutput, fileOutput))

	// 设置日志格式
	common.Log.SetFormatter(&logrus.JSONFormatter{})

	return
}
