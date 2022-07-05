package db

import (
	"fmt"

	"github.com/allanassis/reddere/src/observability/logging"
	"go.mongodb.org/mongo-driver/mongo"
)

func logError(msg string, err error, logger *logging.Logger, db *Database, fields ...logging.Field) {
	logErr := logging.String("error", err.Error())

	newFields := append(fields, logErr)

	if mongo.IsTimeout(err) {
		timeoutField := logging.String("timeout", db.timeout.String())
		timeoutFields := append(newFields, timeoutField)

		newMsg := fmt.Sprintf("Database timeout error: %s", msg)
		logger.Error(newMsg, timeoutFields...)
		return

	} else if mongo.IsNetworkError(err) {
		newMsg := fmt.Sprintf("Database network error: %s", msg)
		logger.Error(newMsg, newFields...)
		return

	} else if mongo.IsDuplicateKeyError(err) {
		newMsg := fmt.Sprintf("Database duplicate key error: %s", msg)
		logger.Error(newMsg, newFields...)
		return

	} else {
		newMsg := fmt.Sprintf("Database unexpected error: %s", msg)
		logger.Error(newMsg, newFields...)
		return
	}
}
