package main

import (
    "fmt"
    "io"
    "os"
	"log"
	"time"
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

//TEST2
type Lead struct {
	lead_id uint32 `json:"lead_id"`
    valid bool `json:"valid"`	
    firstname string `json:"firstname"`	
    lastname string `json:"lastname"`	
    email sql.NullString `json:"email"`	
    phone1 sql.NullString `json:"phone1"`	
	city sql.NullString `json:"city"`	
	state sql.NullString `json:"state"`	
	zipcode sql.NullString `json:"zipcode"`	
	county sql.NullString `json:"county"`	
	contested sql.NullString `json:"contested"`	
	rejected_reason sql.NullString `json:"rejected_reason"`	
	lead_type sql.NullString `json:"contested"`	
	price float32 `json:"amount"`	
	cost float32 `json:"amount"`	
	status sql.NullString `json:"contested"`	
	direction sql.NullString `json:"contested"`	
	tcpa_opted_in sql.NullString `json:"contested"`	
	subid sql.NullString `json:"contested"`	
	comments sql.NullString `json:"contested"`	
	appointment sql.NullString `json:"appointment"`	
    lead_created sql.NullString `json:"lead_created"`	
    practice sql.NullString `json:"practice"`	
    sourcename sql.NullString `json:"sourcename"`	
    trans_id uint32 `json:"id"`	
    trans_created sql.NullString `json:"trans_created"`	
    advertiser_id uint32 `json:"advertiser_id"`	
    amount float32 `json:"amount"`	
    new_balance float32 `json:"new_balance"`	
    transaction_type sql.NullString `json:"transaction_type"`	
    partner_type sql.NullString `json:"partner_type"`	
    firm sql.NullString `json:"firm"`	
}

func main() {
	var records []Record
	current := time.Now()

	//TEST1
	//inputfile
	//var csvFile, err = os.Open("ta/data/ta3.csv")//full records
	//var csvFile, err = os.Open("ta/data/ta5.csv")//32 records
	var csvFile, err = os.Open("ta/data/ta4.csv")//few records

	//logs
	var logfile = "ta/logs/log_" + current.Format("2006-01-02")+".txt"
	conn.DeleteFile(&logfile)

	fp, logerr := os.OpenFile(logfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)//appends to logfile
	if logerr != nil {
		panic(logerr)
	}
	log.SetOutput(fp)

	//output
	var output string = "ta/results/result_"+current.Format("2006-01-02")+".csv"
	if err != nil {
		fmt.Println("Bad. Cannot open file:",err)
	}

	//body
	header := "Starting TA3" 
	fmt.Printf("%v\n", header)
	log.Printf("%v\n", header)

	reader := csv.NewReader(bufio.NewReader(csvFile))	
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

		//fmt.Printf("%v. %v\n",idx, line)//TEST
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

	var data = [][]string{{"Lead Source","Status","Date Added","Last Action Note","First Name","Last Name","Mobile Phone","Home Phone","Email","Lead ID", "Practice", "Lead Source", "Transaction", "Transaction Date", "Payout", 
						   "firm","Valid, "}}

    fmt.Println("-------")
	var max = len(records)
	var cnt = 1
    for i := range records {

		var phones []string

    	fmt.Printf("%v/%v. %v %v, email: %v, phone: %v, mobile: %v\n", cnt, max, records[i].first, records[i].last, records[i].email, records[i].phone, records[i].mobile)

		conn.ConvertsPhone(&phones, records[i].phone)
		conn.ConvertsPhone(&phones, records[i].mobile)

		if len(phones) == 0 && records[i].email == "" { 
			continue
		}

		var results *sql.Rows = conn.SelectRecordfunc(db, phones, records[i].email)

		for results.Next() {//iterate over the rows
	
			var tag Lead 

			//TEST3
			//reads the columns in each row into variable lead with rows.Scan().
			err := results.Scan(&tag.lead_id, &tag.valid, &tag.firstname, &tag.lastname, &tag.email, &tag.phone1, &tag.city, &tag.state, &tag.zipcode, &tag.county, &tag.contested, &tag.rejected_reason, 
								&tag.lead_type, &tag.price, &tag.cost, &tag.status, &tag.direction, &tag.tcpa_opted_in, &tag.subid, &tag.appointment, &tag.comments, &tag.lead_created, &tag.practice, 
								&tag.sourcename, &tag.trans_id, &tag.trans_created, &tag.advertiser_id, &tag.amount, &tag.new_balance, &tag.transaction_type, &tag.partner_type, &tag.firm)

			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

	var data = [][]string{{"Lead Source","Status","Date Added","Last Action Note","First Name","Last Name","Mobile Phone","Home Phone","Email","Lead ID", "Practice", "Lead Source", "Transaction", "Transaction Date", "Payout", 
						   "firm","Valid","Lead Status",""}}
			data = append(data, []string{records[i].leadsource, records[i].status, records[i].dateadded, records[i].lastaction, records[i].first, records[i].last, records[i].mobile, records[i].phone, records[i].email, 
						  fmt.Sprint(tag.lead_id), fmt.Sprint(tag.valid), tag.contested.String, tag.rejected_reason.String, tag.lead_type.String, fmt.Sprint(tag.price), fmt.Sprint(tag.cost), tag.status.String, 
						  tag.city.String, tag.state.String, tag.zipcode.String, tag.county.String, tag.direction.String, tag.tcpa_opted_in.String, tag.comments.String, tag.lead_created.String, 
						  tag.practice.String, tag.sourcename.String, tag.subid.String, tag.appointment.String, fmt.Sprint(tag.trans_id), tag.trans_created.String, fmt.Sprint(tag.amount), tag.firm.String})

		}//for

		cnt+=1

    }//for
    
	//output records
	
	conn.Writefile(&output, data)
    
}//func

//misc


