package helpers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/labstack/echo"
	"time"
)

var Logger LogConfig

type LogConfig struct {
	level       string
	logLevelMap map[string]int
}

func InitLogger(printLevelUpto string) {
	Logger = LogConfig{
		logLevelMap: map[string]int{
			"info":    0,
			"warning": 1,
			"error":   2,
			"fatal":   3,
		},
		level: printLevelUpto,
	}
}

func (l *LogConfig) messageMaker(ctx echo.Context, message string) (formattedMsg string) {
	currentTime := time.Now()
	var serviceName = " - "
	var requestID = " - "
	if ctx == nil {
		formattedMsg = fmt.Sprintf("%s | %s | %s | %s", currentTime.Format(time.RFC1123),
			requestID, serviceName, message)
		return
	}
	serviceNameVal, ok := ctx.Get("serviceName").(string)
	if ok {
		serviceName = serviceNameVal
	}
	requestIDVal, ok := ctx.Get("requestID").(string)
	if ok {
		requestID = requestIDVal
	}
	formattedMsg = fmt.Sprintf("%s | %s | %s | %s", currentTime.Format(time.RFC1123),
		requestID, serviceName, message)
	return
}

func (l *LogConfig) Info(ctx echo.Context, message string) {
	if l.logLevelMap[l.level] <= l.logLevelMap["info"] {
		color.Cyan(l.messageMaker(ctx, message))
	}
}

func (l *LogConfig) Error(ctx echo.Context, message string, err error) {
	if l.logLevelMap[l.level] <= l.logLevelMap["error"] {
		color.Red(l.messageMaker(ctx, message+err.Error()))
	}
}
func (l *LogConfig) Warn(ctx echo.Context, message string) {
	if l.logLevelMap[l.level] <= l.logLevelMap["warn"] {
		color.Yellow(l.messageMaker(ctx, message))
	}
}
