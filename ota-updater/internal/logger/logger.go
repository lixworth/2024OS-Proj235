package logger

import (
	"github.com/nyancatda/AyaLog/v2"
)

var (
	Logger *AyaLog.Log
)

func InitLogger() {
	Logger = AyaLog.NewLog()
	Logger.ColorPrint = true
	Logger.WriteFile = false
}

func GetLogger() *AyaLog.Log {
	return Logger
}
