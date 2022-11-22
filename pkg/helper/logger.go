package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"os"
	"path"
)

func InitLogger(execDir, webLogPath, mode string) (sl *slog.SugaredLogger, err error) {
	// 检测文件夹是否存在
	logPath := path.Join(execDir, webLogPath)
	if !PathExists(logPath) {
		_ = os.Mkdir(logPath, os.ModePerm)
	}
	logName := path.Join(logPath, "log")
	file, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	// debug阶段记录debug以及以上错误
	// release阶段记录info及以上错误
	var logLevel slog.Level
	if mode == gin.DebugMode {
		logLevel = slog.DebugLevel
	} else {
		logLevel = slog.InfoLevel
	}
	return slog.NewJSONSugared(file, logLevel), nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
