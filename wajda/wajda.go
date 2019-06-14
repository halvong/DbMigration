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
    conn "DbMigration/wajda/db"
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
    firstname sql.NullString `json:"firstname"`	
    lastname sql.NullString `json:"lastname"`	
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
    firm_id uint32 `json:"firm_id"`	
    firm_first sql.NullString `json:"firm_first"`	
    firm_last sql.NullString `json:"firm_last"`	
    firm sql.NullString `json:"firm"`	
}

const WHERE string = "ALL"

func main() {

	var records []Record //input data from read csv file
	var output string //output
	var	data = [][]string{{"Lead ID","Valid","Lead Status","Contested","Rejected Reason","Lead Type", "Price","Cost","City","State","Zipcode","County","Direction","TCPA","Comments","Lead Created","Practice",
							"Lead Source", "Subid","Appointment","Transaction Id","Transaction Created", "Transaction Amount", "Firm Id (Advertiser)","Firm First Name (Advertiser)","Firm Last Name (Advertiser)",
							"Firm (Advertiser)" }}

	//logs
	current := time.Now()
	var logfile = "wajda/logs/log_" + current.Format("2006-01-02")+".txt"
	conn.DeleteFile(&logfile)

	//db connection
	var db *sql.DB = conn.Connectfunc("LIVE")
    defer db.Close()

	//title
	header := "Starting TA3" 
	fmt.Printf("%v\n", header)
	log.Printf("%v\n", header)

	if WHERE == "ALL" {

		output = "wajda/results/leads_"+current.Format("2006-01-02")+".csv"

		
	} else {

		output = "wajda/results/wajda_"+current.Format("2006-01-02")+".csv"

		var results *sql.Rows = conn.SelectFromAdvertiserfunc(db, 2125, 2019)
	
		for results.Next() {//iterate over the rows

			var tag Lead 

			//reads the columns in each row into variable lead with rows.Scan().
			err := results.Scan(&tag.lead_id, &tag.valid, &tag.firstname, &tag.lastname, &tag.email, &tag.phone1, &tag.city, &tag.state, &tag.zipcode, &tag.county, &tag.contested, &tag.rejected_reason, 
								&tag.lead_type, &tag.price, &tag.cost, &tag.status, &tag.direction, &tag.tcpa_opted_in, &tag.subid, &tag.appointment, &tag.comments, &tag.lead_created, &tag.practice, 
								&tag.sourcename, &tag.trans_id, &tag.trans_created, &tag.advertiser_id, &tag.amount, &tag.new_balance, &tag.transaction_type, &tag.partner_type, &tag.firm_id, &tag.firm_first, &tag.firm_last, 
								&tag.firm)

			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

			data = append(data, []string{fmt.Sprint(tag.lead_id), fmt.Sprint(tag.valid), tag.status.String, tag.contested.String, tag.rejected_reason.String, tag.lead_type.String, fmt.Sprint(tag.price), fmt.Sprint(tag.cost), 
						  tag.city.String, tag.state.String, tag.zipcode.String, tag.county.String, tag.direction.String, tag.tcpa_opted_in.String, tag.comments.String, tag.lead_created.String, 
						  tag.practice.String, tag.sourcename.String, tag.subid.String, tag.appointment.String, fmt.Sprint(tag.trans_id), tag.trans_created.String, fmt.Sprint(tag.amount), fmt.Sprint(tag.firm_id), tag.firm_first.String, tag.firm_last.String, tag.firm.String})		  

		}//for
	
	}
	
	//output records
	if len(data) > 0 {
		conn.Writefile(&output, data)
	}	
    
}//func


