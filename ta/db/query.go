package db

import (
	"fmt"
	_ "strconv"
	"database/sql"
	_ "strings"
)



func SelectRecordfunc(db *sql.DB, phones []string, email string) *sql.Rows {
	var str string
	var display string

	for _, value := range phones {
		str += value + "," 
		//str += strconv.Itoa(value) + ","
	}	

	var sql string = "SELECT lead.id, lead.firstname, lead.lastname, lead.email, lead.phone1, source.name, trans.id, trans.advertiser_id, trans.amount, trans.new_balance, trans.transaction_type," 
	sql += " trans.partner_type, adv.firm FROM attorney_lead lead" 
	sql += " INNER JOIN attorney_leadsource source ON lead.source_id = source.id"
	sql += " INNER JOIN attorney_transaction trans ON lead.id = trans.lead_id" 
	sql += " INNER JOIN attorney_advertiser adv ON adv.id = trans.advertiser_id" 

	if len(phones) > 0 {
		sql += " WHERE lead.email = ? OR lead.phone1 IN (" + str[:len(str)-1] + ")" 
		display = " WHERE lead.email = ? OR lead.phone1 IN (" + str[:len(str)-1] + ")" 
	} else {
		sql += " WHERE lead.email = ?"
		display = " WHERE lead.email = ?"
	}

	sql += " ORDER BY lead.created DESC"
	fmt.Printf("\tsql: %v\n", display)
	
	stmt, err := db.Prepare(sql)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	results, err := stmt.Query(email)
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