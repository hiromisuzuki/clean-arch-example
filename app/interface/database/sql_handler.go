package database

//SQLHandler SQLHandler interface
type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

//Result Result interface
type Result interface {
	LastInsertID() (int64, error)
	RowsAffected() (int64, error)
}

//Row Row interface
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
