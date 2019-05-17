package db

import (
	"os"
	"log"
	"fmt"
	"regexp"
	"database/sql"
)

var re = regexp.MustCompile(`^\((\d{3})\)\s(\d{3})-(\d{4})`)

func CheckError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func ConvertsPhone(phones *[]string, phone string) {

		var res string

		match := re.FindStringSubmatch(phone)
		if len(match) == 4 {
			res = fmt.Sprintf("1%v%v%v", match[1], match[2], match[3])
		}

		if res != "" {
			*phones = append(*phones, res)
		}
		//return res
}

func DeleteFile(infile_ptr *string) bool {
    // deletes file, TEST3
    _, err := os.Stat(*infile_ptr)

	if err == nil {

		err := os.Remove(*infile_ptr)

		if err != nil {
			fmt.Printf("\tFile %v failed to delete.", *infile_ptr)
			return false
		}
		return true
	}

	return false
}

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