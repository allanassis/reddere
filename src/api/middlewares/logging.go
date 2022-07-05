package middlewares

import (
	"time"

	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequestLogger(logger *logging.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			eventID, err := uuid.NewRandom()
			if err != nil {
				panic(err)
			}

			req := ctx.Request()
			requestId := ctx.Request().Header.Get(echo.HeaderXRequestID)

			logger.UpdateFiels(logging.String("eventID", eventID.String()))
			logger.UpdateFiels(logging.String("requestID", requestId))

			fields := []logging.Field{
				logging.String("remoteIp", ctx.RealIP()),
				logging.String("host", req.Host),
				logging.String("method", req.Method),
				logging.String("path", req.RequestURI),
				logging.String("userAgent", req.UserAgent()),
				logging.Int64("size", req.ContentLength),
			}

			logger.Info("Request to the api", fields...)

			err = next(ctx)
			if err != nil {
				return err
			}

			return nil
		}
	}
}

func ResponseLogger(logger *logging.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			start := time.Now()

			res := ctx.Response()
			err := next(ctx)
			if err != nil {
				return err
			}

			fields := []logging.Field{
				logging.String("latency", time.Since(start).String()),
				logging.Int("status", res.Status),
				logging.Int64("size", res.Size),
			}

			logger.Info("Response from the api", fields...)
			logger.CleanUpFields()
			return nil
		}
	}
}
