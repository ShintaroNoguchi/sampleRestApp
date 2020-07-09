package myerrors

// DB is struct for database error
type DB struct {
	Message string
}

// NewDB for database error
func NewDB(str string) DB {
	return DB{str}
}

func (e DB) Error() string {
	return e.Message
}
