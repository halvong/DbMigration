package db

import (
	"fmt"
	"strconv"
	"database/sql"
	_ "strings"
)

func SelectLeadIDfunc(db *sql.DB, ids []int, sourceid int) *sql.Rows {

	var str string
	for _, value := range ids {
		str += strconv.Itoa(value) + ","
	}	

	var sql string = "SELECT id, firstname, lastname, email, contested, rejected_reason, status, price, area_id, source_id, advertisement_id, upright_law_api_response, created FROM attorney_lead WHERE advertisement_id = ? AND id IN (" + str[:len(str)-1] + ") ORDER BY created DESC"
	fmt.Printf("sql: %v\n", sql)
	
	stmt, err := db.Prepare(sql)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	results, err := stmt.Query(sourceid)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	return results
}

func SelectLeadfunc(db *sql.DB, sourceid int) *sql.Rows {
	
	stmt, err := db.Prepare("SELECT id, firstname, lastname, email, contested, rejected_reason, status, price, area_id, source_id, advertisement_id, upright_law_api_response, created from attorney_lead WHERE advertisement_id = ? ORDER BY created DESC LIMIT 10")

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	results, err := stmt.Query(sourceid)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	return results
}