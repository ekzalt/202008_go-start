package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // https://github.com/go-sql-driver/mysql
)

// User is the db entity
type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func main() {
	driverName := "mysql"
	dataSourceName := "root:root@tcp(127.0.0.1:8889)/dbname" // "user:password@/dbname"

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		panic(err)
	}

	// close db connection in the end
	defer db.Close()

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	/*
		// insert data
		rows, err := db.Query("INSER INTO `users` (`name`, `age`) VALUES ('Alex', 25)")

		if err != nil {
			panic(err)
		}

		// release connection from the pool
		defer rows.Close()
	*/

	// select data
	rows, err := db.Query("SELECT * FROM `users`")

	if err != nil {
		panic(err)
	}

	// release connection from the pool
	defer rows.Close()

	for rows.Next() {
		user := User{}
		scanErr := rows.Scan(&user)

		if scanErr != nil {
			panic(scanErr)
		}

		fmt.Println(user)
	}
}
