package db

import (
	"database/sql"
)

func SelectTransactionfunc(db *sql.DB, email string, sourceid int) *sql.Rows {
	
	stmt, err := db.Prepare("SELECT l.id, l.firstname, l.lastname, l.rejected_reason, l.contested, l.status FROM attorney_lead l LEFT JOIN attorney_transaction t ON (l.id = t.lead_id) WHERE l.email = ? AND source_id= ? ORDER BY l.created DESC limit 5")

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	results, err := stmt.Query(email, sourceid)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	return results
}