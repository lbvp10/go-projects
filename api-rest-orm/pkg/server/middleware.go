package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"orm"
	"orm/pkg/logger"
	"os"
	"time"
)

// MiddlewareRequestLogs middleware que intercepta todas las peticiones que llegan a la API y las envia al log
type MiddlewareRequestLogs interface {
	LoggerApiRequest(ctx *fiber.Ctx) error
}

type logsMiddleware struct {
	log logger.Logger
}

func NewLogsMiddleware(log logger.Logger) MiddlewareRequestLogs {
	return &logsMiddleware{log: log}
}

func (s *logsMiddleware) LoggerApiRequest(ctx *fiber.Ctx) error {
	t := time.Now()
	fields := map[string]interface{}{
		"IP":        ctx.IP(),
		"URL":       ctx.OriginalURL(),
		"StartTime": t.Format(api_rest_orm.DateFormat),
		"Method":    ctx.Method(),
		"RequestId": ctx.Locals("requestid").(string),
	}

	result := ctx.Next()

	if GetConfigLogsRequest() {
		fields["RequestHeader"] = ctx.GetReqHeaders()
		fields["RequestParam"] = ctx.Request().URI().QueryArgs().String()
		fields["ResponseBody"] = string(ctx.Response().Body()[:])
	}

	fields["Status"] = ctx.Response().StatusCode()
	fields["EndTime"] = time.Now().Format(api_rest_orm.DateFormat)
	fields["Duration"] = time.Since(t).Milliseconds()

	s.log.Fields(fmt.Sprintf("%s %s", ctx.Method(), fields["URL"]), fields)

	return result
}
func GetConfigLogsRequest() bool {
	if os.Getenv("LOG_FULL") == "false" {
		return false
	}
	return true
}
