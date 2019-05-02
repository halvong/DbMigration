package main

import (
	"fmt"
    "database/sql"
	_ "reflect"
	"time"
	conn "DbMigration/mysql/db"
)

type Transaction struct {
	id   int    `json:"id"`
    firstname string `json:"firstname"`	
    lastname string `json:"lastname"`	
    rejected_reason sql.NullString `json:"rejected_reason"`	
    contested sql.NullString `json:"contested"`	
    status sql.NullString `json:"status"`	
}

type Lead struct {
	id   int    `json:"id"`
    firstname string `json:"firstname"`	
    lastname string `json:"lastname"`	
    email sql.NullString `json:"email"`	
    contested sql.NullString `json:"contested"`	
    rejected_reason sql.NullString `json:"rejected_reason"`	
    status sql.NullString `json:"status"`	
    price float32 `json:"price"`	
    area_id uint8 `json:"area_id"`	
    source_id uint8 `json:"source_id"`	
    advertisement_id uint16 `json:"advertisement_id"`	
    upright_law_api_response sql.NullString `json:"upright_law_api_response"`	
	created time.Time `json:"created"`	
}

func main() {

    fmt.Println("Go MySQL")
	var db *sql.DB = conn.Connectfunc("LIVE")
    defer db.Close()

	ids := []int{1230299, 1230304, 1230252, 1230050, 1228737, 1228130, 1227889, 1230743}

	var results *sql.Rows = conn.SelectLeadIDfunc(db, ids, 6404) 

	var idx int = 1
	for results.Next() {

		var tag Lead

		err := results.Scan(&tag.id, &tag.firstname, &tag.lastname, &tag.email, &tag.contested, &tag.rejected_reason, &tag.status, &tag.price, 
							&tag.area_id, &tag.source_id, &tag.advertisement_id, &tag.upright_law_api_response, &tag.created)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		fmt.Printf("%d. id: %v, firstname: %v, lastname: %v, email: %v, contested: %v, rejected_reason: %v, status: %v,\n\tprice: %v,area_id: %v, source_id: %v, advertisement_id: %v,upright_law_api_response: %v, created: %v+\n\n", idx, tag.id, tag.firstname, tag.lastname, tag.email, tag.contested, tag.rejected_reason, 
			tag.status,
			tag.price,
			tag.area_id, 
			tag.source_id, 
			tag.advertisement_id, tag.upright_law_api_response, tag.created)

		idx += 1

	}//for

}

	/*
	customers := []string { "ashley9760@gmail.com", "dle621092@gmail.com", "floroarlene@gmail.com", "kbjones9210@gmail.com", "jerome.pipkins1426@gmail.com" }

	var idx int = 0
    fmt.Println("line, lead id, firstname, lastname, rejected reason, contested, status")
	for ii := range customers {

		results = conn.SelectTransactionfunc(db, customers[ii], 171)

		for results.Next() {

		var tag Transaction
			// for each row, scan the result into our tag composite object
			//err := results.Scan(&tag.id, &tag.firstname, &tag.lastname, &tag.rejected_reason, &tag.contested, &tag.status, &tag.price, &tag.area_id, &tag.source_id, &tag.advertisement_id, &tag.created)
			err := results.Scan(&tag.id, &tag.firstname, &tag.lastname, &tag.rejected_reason, &tag.contested, &tag.status)

			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			fmt.Printf("%d, %v, %v, %v, %v, %v, %v\n", idx + 1, tag.id, tag.firstname, tag.lastname, tag.rejected_reason, tag.contested, tag.status)
			idx++ 
		}

	}//for
	*/


