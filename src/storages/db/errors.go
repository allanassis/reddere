package db

type DataBaseError struct {
	Msg DbError
	Err error
}

func (db DataBaseError) Error() string {
	return string(db.Msg)
}

func (db DataBaseError) InternalError() error {
	return db.Err
}

type DbError string

func newDatabaseError(errCode DbError, err error) DataBaseError {
	return DataBaseError{
		Msg: errCode,
		Err: err,
	}
}

const (
	DB_HEALTHCHECK_ERROR   DbError = "Error when ping to database"
	DB_BIND_ERROR          DbError = "Error when binding database document to instance"
	DB_DELETE_ONE_ERROR    DbError = "Error when deleting item in database"
	DB_INVALID_ID_ERROR    DbError = "Error when tried to parse invalid id"
	DB_FIND_ONE_ERROR      DbError = "Error when tried to find document in database"
	DB_INSER_ONE_ERROR     DbError = "Error when tried to insert document into database"
	DB_CREATE_CLIENT_ERROR DbError = "Error when tried to create database client"
)
