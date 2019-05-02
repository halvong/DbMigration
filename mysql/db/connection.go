package db

import (
	"fmt"
    "database/sql"
	_ "github.com/go-sql-driver/mysql" // _ registers for initialization (init) as database driver
)

func Connectfunc(machine string) *sql.DB {

	var db *sql.DB
	var err error


	// Open up our database connection.
	if machine == "LIVE" {
		db, err = sql.Open("mysql", "mainuser:IWon'tGrowUp@tcp(127.0.0.1:3306)/web_main_live?parseTime=true")
	} else {
		db, err = sql.Open("mysql", "mainuser:IWon'tGrowUp@tcp(127.0.0.1:3306)/web_main_qa?parseTime=true")
	}

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

	fmt.Println("Connection Successful")

	return db
}
