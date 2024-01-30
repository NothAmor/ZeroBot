package safe

import (
	"runtime/debug"

	"github.com/NothAmor/ZeroBot/core/common"
)

func Go(goroutine func()) {
	GoWithRecover(goroutine, func(err interface{}) {
		common.Log.Errorf("Error in Go routine: %s\nStack: %s", err, debug.Stack())
	})
}

func GoWithRecover(goroutine func(), customRecover func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				customRecover(err)
			}
		}()
		goroutine()
	}()
}
