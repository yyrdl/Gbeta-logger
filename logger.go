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
	green     = color.New(color.FgGreen).SprintFunc()
	yellow    = color.New(color.FgYellow).SprintFunc()
	red       = color.New(color.FgRed).SprintFunc()
	blue      = color.New(color.FgBlue).SprintFunc()
	cyan      = color.New(color.FgCyan).SprintFunc()
	hiwhite   = color.New(color.FgHiWhite).SprintFunc()
	himagenta = color.New(color.FgHiMagenta).SprintFunc()
)

var statusColor = func(status int) string {
	status_msg := fmt.Sprintf("%d %s", status, http.StatusText(status))
	switch {
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
			return cyan(status_msg)
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
	switch {
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

var logger = log.New(os.Stdout, green("[Gebeta]"), log.Ldate|log.Ltime)

func Log(serve_http gbeta.ServeHTTPFunc) gbeta.ServeHTTPFunc {
	return func(res gbeta.Res, req gbeta.Req) {
		start := time.Now()
		serve_http(res, req)
		logger.Printf("  %s  %s  %s  %v", himagenta(req.Method), green(req.URL.Path), statusColor(res.Code()), timeColor(time.Since(start)))
		return
	}
}
