package loghook

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

type LineHook struct{}

func (h *LineHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LineHook) Fire(entry *logrus.Entry) error {
	pc, file, line, ok := runtime.Caller(8)
	if !ok {
		return nil
	}

	funcName := runtime.FuncForPC(pc).Name()
	entry.Data["caller"] = fmt.Sprintf("%s:%d (%s)", file, line, funcName)

	return nil
}
