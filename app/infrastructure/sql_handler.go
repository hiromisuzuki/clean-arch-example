package infrastructure

import (
	"database/sql"
	"fmt"
	"net/url"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/hiromisuzuki/clean-arch-example/app/interface/database"
)

//SQLConfig SQLConfig
type SQLConfig struct {
	DBMS     string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Options  *url.Values
}

//DSN DSN
func (c *SQLConfig) DSN() (dsn string) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
	dsn = fmt.Sprintf("%s?%s", conn, c.Options.Encode())
	return dsn
}

//SQLHandler SQLHandler
type SQLHandler struct{ Conn *sql.DB }

//NewSQLHandler NewSQLHandler
func NewSQLHandler(config *SQLConfig) *SQLHandler {
	//TODO:define configure
	conn, err := sql.Open(config.DBMS, config.DSN())
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

//Execute Execute
func (handler *SQLHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

//Query Query
func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	row := SQLRow{}
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return row, err
	}
	row.Rows = rows
	return row, nil
}

//SQLResult SQLResult
type SQLResult struct{ Result sql.Result }

//LastInsertID LastInsertID
func (r SQLResult) LastInsertID() (int64, error) {
	return r.Result.LastInsertId()
}

//RowsAffected RowsAffected
func (r SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

//SQLRow SQLRow
type SQLRow struct{ Rows *sql.Rows }

//Scan Scan
func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

//Next Next
func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

//Close Close
func (r SQLRow) Close() error {
	return r.Rows.Close()
}
