package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

const (
	blackText = "\033[0;30m"
	orange    = "\033[48;2;255;165;0m"
	cyan      = "\033[46m"
	red       = "\033[41m"
	green     = "\033[42m"
	blue      = "\033[44m"
	white     = "\033[47m"
	yellow    = "\033[43m"
	black     = "\033[40m"
	reset     = "\033[0m"
)

func SetupLogOutput() {
	f, _ := os.Create("./log/output.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s%s%s%s %s[%s]%s %s%s%s%s %s%s%s%s %s%s%d%s %s%s%s%s\n",
			blackText, white, param.ClientIP, reset,
			black, param.TimeStamp.Format(time.TimeOnly), reset,
			blackText, methodColor(param.Method), param.Method, reset,
			blackText, white, param.Path, reset,
			blackText, codeColor(param.StatusCode), param.StatusCode, reset,
			blackText, white, param.Latency, reset,
		)
	})
}

func methodColor(method string) string {
	switch method {
	case "GET":
		return green
	case "POST":
		return orange
	case "PUT":
		return cyan
	case "DELETE":
		return red
	default:
		return white
	}
}

func codeColor(code int) string {
	switch code / 100 {
	case 1:
		return blue
	case 2:
		return green
	case 3:
		return yellow
	case 4:
		return red
	case 5:
		return red
	default:
		return white
	}
}
