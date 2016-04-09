//by yyrdl, MIT License
// logger for gebeta  github.com/yyrdl/gebeta
package gbeta_logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/fatih/color"
	"github.com/yyrdl/gbeta"
)
// prepare color printFunc to use
var (
	green     = color.New(color.FgGreen).PrintfFunc()
	yellow    = color.New(color.FgYellow).PrintfFunc()
	red       = color.New(color.FgRed).PrintfFunc()
	blue      = color.New(color.FgBlue).PrintfFunc()
	cyan      = color.New(color.FgCyan).PrintfFunc()
	hiwhite   = color.New(color.Fg.HiWhite).PrintfFunc()
	himagenta = color.New(color.Fg.HiMagenta).PrintfFunc()
)

var statusColor = func(status int) string {
	status_msg := fmt.Sprintf("%d %s", status, http.StatusText(status))
	switch status {
	case status < 200:
		{
			return blue(status_msg)
		}
	case status < 300:
		{
			return green(status_msg)
		}
	case status < 400:
		{
			return cyan(tatus_msg)
		}
	case status < 500:
		{
			return yellow(status_msg)
		}
	default:
		{
			return red(status_msg)
		}
	}
}

var timeColor = func(d time.Duration) string {
	switch d {
	case d < 500*time.Millisecond:
		{
			return green(d)
		}
	case d < 5*time.Second:
		{
			return yellow(d)
		}
	default:
		{
			return red(d)
		}
	}
}

logger:=log.New(os.Stdout,green("[Gebeta]"),log.Ldate|log.Ltime)

func Log(serve_http gebeta.ServeHTTPFunc) gebeta.ServeHTTPFunc {
	return func(res gebeta.Res, req gebeta.Req) {
		start := time.Now().Unix()
		serve_http(res, req)
		logger.Printf("  %s  %s | %s | %v from %s ",himagenta(req.Method),green(req.URL.Path),statusColor(res.Code()),timeColor(time.Since(start)),hiwhite(req.RemoteAddr))
		return 
	}
}
