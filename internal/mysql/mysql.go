package mysql

import (
	"database/sql"
	"fmt"
	"net/url"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New(u url.URL) (*sql.DB, error) {
	uri := fmt.Sprintf("%s@tcp(%s)/", u.User.String(), u.Host)
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return db, err
	}

	maxProcs := runtime.GOMAXPROCS(0)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(maxProcs)
	db.SetMaxIdleConns(maxProcs)

	return db, err
}
