package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"stock_fetchernstore/internal/config"
	"stock_fetchernstore/internal/mysql"
	"time"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.Load()
	db := getDb()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	dbname := viper.GetString("database.mysql.db")
	Exec(db, fmt.Sprintf(`

	CREATE DATABASE IF NOT EXISTS %s;

	`, dbname))

	Exec(db, fmt.Sprintf(`
	
	USE %s;

	`, dbname))

	Exec(db, fmt.Sprintf(`
	
	CREATE TABLE IF NOT EXISTS stock (
		id INT NOT NULL PRIMARY KEY,
		symbol VARCHAR(20),
		close FLOAT,
		high FLOAT,
		low FLOAT,
		open FLOAT,
		dt DATETIME
	);

	`))

	// Exec(dbc

}

func Exec(db *sql.DB, q string, params ...any) sql.Result {
	timeout := viper.GetInt64("database.mysql.timeout")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	var result sql.Result
	var err error

	if len(params) > 0 {
		result, err = db.ExecContext(ctx, q)
	} else {
		result, err = db.ExecContext(ctx, q, params...)
	}

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDb() *sql.DB {
	user := viper.GetString("database.mysql.username")
	passwd := viper.GetString("database.mysql.password")
	host := viper.GetString("database.mysql.host")

	u := url.URL{
		Host: host,
		User: url.UserPassword(user, passwd),
	}

	db, err := mysql.New(u)
	if err != nil {
		os.Exit(1)
	}

	return db
}
