package middlewares

import (
	"time"

	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/labstack/echo/v4"
)

func Logger(log *logging.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			start := time.Now()

			err := next(ctx)
			if err != nil {
				return err
			}

			req := ctx.Request()
			res := ctx.Response()

			status := res.Status

			fields := []logging.Field{
				logging.String("request_id", req.Header.Get(echo.HeaderXRequestID)),
				logging.String("remote_ip", ctx.RealIP()),
				logging.String("latency", time.Since(start).String()),
				logging.String("host", req.Host),
				logging.String("method", req.Method),
				logging.String("path", req.RequestURI),
				logging.String("user_agent", req.UserAgent()),
				logging.Int64("size", res.Size),
				logging.Int("status", status),
			}

			log.Info("Request to the api", fields...)

			return nil
		}
	}
}
