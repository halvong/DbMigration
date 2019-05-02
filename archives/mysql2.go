package main

import (
	"fmt"
	"reflect"
    "database/sql"
	"time"
    _ "github.com/go-sql-driver/mysql" // _ registers for initialization (init) as database driver
)

type Tag struct {
	id   int    `json:"id"`
    firstname string `json:"firstname"`	
    lastname string `json:"lastname"`	
    contested string `json:"contested"`	
    rejected_reason string `json:"rejected_reason"`	
    status string `json:"status"`	
    price float32 `json:"price"`	
    area_id uint8 `json:"area_id"`	
    source_id uint32 `json:"source_id"`	
    advertisement_id uint16 `json:"advertisement_id"`	
	created time.Time `json:"created"`	
	transaction_type string `json:"transaction_type"`	
partner_type
string `json:""`	

advertiser_id
uint32 `json:""`	

lead_source_id
uint32 `json:""`	

amount
float32 `json:""`	

new_balance
float32 `json:""`	

lead_id
uint32 `json:""`	

created
time.Time `json:"created"`	
}


func connectfunc() *sql.DB {

    // Open up our database connection.
    db, err := sql.Open("mysql", "mainuser:IWon'tGrowUp@tcp(127.0.0.1:3306)/web_main_live")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

	fmt.Println("Connection Successful")
	fmt.Println(reflect.TypeOf(db))
	
    //defer db.Close()

	return db
}


func main() {

    fmt.Println("Go MySQL")

	var db *sql.DB

	db = connectfunc()

	fmt.Println(reflect.TypeOf(db))

	results, err := db.Query("SELECT l.id, l.firstname, l.lastname, l.contested, l.rejected_reason, l.status, l.price, l.area_id, l.source_id, l.advertisement_id, l.created, t.transaction_type, t.partner_type, t.advertiser_id, t.lead_source_id, t.amount, t.new_balance, t.lead_id, t.created AS 'Transaction created' from attorney_lead l join attorney_transaction t on (l.id = t.lead_id) WHERE l.firstname LIKE '%Darious%' and l.lastname LIKE '%Walker%' order by l.created DESC limit 5")

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {

        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.id, &tag.firstname)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        fmt.Printf("%v, %v",tag.id,tag.firstname)
    }

    db.Close()
}