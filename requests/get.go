package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"encoding/json"
	"time"
	"log"
	"net/url"
)

type Url struct {
	api_key string	
	zipcode string
	area string
}

func main() {
	//MakeRequest()
	MakeZipcodesRequest()
}

func Request(addr string) {
	//log.Println(addr)

	resp, err := http.Get(addr)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func MakeZipcodesRequest() {

	var zipcodes = []string{"99210", "99211"}

	obj := Url{api_key: "A30F3C4BCE7147C980854343F390AA5E",
			   area: "Auto Accidents",
			}

	var addr string = "https://local.legalzoom.com/apis/ping_form_lead?"
	addr += "api_key="+obj.api_key 
	addr += "&area="+url.QueryEscape(obj.area)

	for idx, zip := range zipcodes {
		final_addr := addr + "&zipcode="+zip
		fmt.Printf("%v. Calling %v\n", idx + 1, final_addr)
		Request(final_addr) 
		time.Sleep(2)
	}
}

func MakeRequest() {

	//addr := "https://local.legalzoom.com/apis/ping_form_lead?api_key=53BF1B7524684A32BCA4EE8BB78A76CF&zipcode=07961&physically_injured=Yes&accepted_terms_of_service=Yes&email=e95cbca0c629a135099fc21372a55ba9&phone=2b05ef0e18280d3659a8aa4013ed84cd&area="
	obj := Url{api_key: "A30F3C4BCE7147C980854343F390AA5E",
			   zipcode: "24014",	
			   area: "Auto Accidents",
			}

	var addr string = "https://local.legalzoom.com/apis/ping_form_lead?"

	addr += "api_key="+obj.api_key 
	addr += "&zipcode="+obj.zipcode
	addr += "&area="+url.QueryEscape(obj.area)

	Request(addr) 

}

