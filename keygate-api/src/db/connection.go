package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	cfg := mysql.Config{
        User:   user,
        Passwd: password,
		AllowNativePasswords: true,
        Net:    "tcp",
        Addr:   host,
        DBName: dbName,
    }

	var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

	pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }


	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database.")
}

func Conn(ctx context.Context) (*sql.Conn, error) {
	return db.Conn(ctx)
}

func GetDB() *sql.DB {
	return db
}