package main

import (
    "fmt"
    "io"
    "os"
    "bufio"
    "encoding/csv"
	_"strconv"
	"database/sql"
    conn "DbMigration/ta/db"
)


type Record struct {
	leadsource string `json:"leadsource"`		
	status string `json:"status"`		
	dateadded string `json:"dateadded"`		
	lastaction string `json:"lastaction"`		
	first string `json:"firstname"`		
	last string `json:"lastname"`		
	mobile string `json:"mobile,omitempty"`		
	phone string `json:"phone,omitempty"`		
	email string `json:"email,omitempty"`		
}

type Lead struct {
	id uint32 `json:"id"`
    firstname string `json:"firstname"`	
    lastname string `json:"lastname"`	
    email sql.NullString `json:"email"`	
    phone1 sql.NullString `json:"phone1"`	
    name sql.NullString `json:"name"`	
    trans_id uint32 `json:"id"`	
    advertiser_id uint32 `json:"advertiser_id"`	
    amount float32 `json:"amount"`	
    new_balance float32 `json:"new_balance"`	
    transaction_type sql.NullString `json:"transaction_type"`	
    partner_type sql.NullString `json:"partner_type"`	
    firm sql.NullString `json:"firm"`	
}

func main() {
	 
	var records []Record

	csvFile, err := os.Open("ta/ta3.csv")

	if err != nil {
		fmt.Println("Bad. Cannot open file:",err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))	
	_ = reader
	
	fmt.Printf("Starting TA3 %v\n", "World")

	var idx int = 0 
	for {
		
		line, err := reader.Read()	

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Bad. Cannot open line:",err)
		}

		if idx == 0 {
			idx += 1
			continue
		} 

		//fmt.Printf("%v. %v\n",idx, line[8])
		records = append(records, Record {
										leadsource: line[0],
										status: line[1],
										dateadded:line[2],
										lastaction:line[3],
										first: line[4],
										last: line[5],
										mobile: line[6],
										phone: line[7],
										email: line[8],
						})

		idx += 1
	}

	var db *sql.DB = conn.Connectfunc("QA")
    defer db.Close()


	//var data [][]string
	var data = [][]string{{"Lead Source","Status","Date Added","Last Action Note","First Name","Last Name","Mobile Phone","Home Phone","Email","Lead ID", "Lead Source"}}

    fmt.Println("-------")
	var max = len(records)
	var cnt = 1
    for i := range records {

		var phones []string

    	fmt.Printf("%v/%v. %v %v, email: %v, phone: %v, mobile: %v\n", cnt, max, records[i].first, records[i].last, records[i].email, records[i].phone, records[i].mobile)

		conn.ConvertsPhone(&phones, records[i].phone)
		conn.ConvertsPhone(&phones, records[i].mobile)

		var results *sql.Rows = conn.SelectRecordfunc(db, phones, records[i].email)

		for results.Next() {

			var tag Lead 

			err := results.Scan(&tag.id, &tag.firstname, &tag.lastname, &tag.email, &tag.phone1, &tag.name, &tag.trans_id, &tag.advertiser_id, &tag.amount, &tag.new_balance, &tag.transaction_type,
								&tag.partner_type, &tag.firm)

			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			data = append(data, []string{records[i].leadsource, records[i].status, records[i].dateadded, records[i].lastaction, records[i].first, records[i].last, records[i].mobile, records[i].phone, records[i].email, fmt.Sprint(tag.id), tag.firm.String})

		}//for

		cnt+=1

    }//for
    
	//write csv file
	var output string = "ta/result.csv"
	delete_bool := conn.DeleteFile(&output)
	if delete_bool {
		fmt.Println("\nDeletes old result.csv successful.")
	}

	file, err := os.Create(output)
    conn.CheckError("Cannot create file", err)
    defer file.Close()
	writer := csv.NewWriter(file)
    defer writer.Flush()
    
	fmt.Println("Writes to result.csv.")
	writer.WriteAll(data)
    
}//func

//misc


