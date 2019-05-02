package main

import (
	"fmt"
	"regexp"
)


func main() {

    fmt.Println("Running misc")

	var title string = "main_live_adwords_geotarget.sql"

	match, _ := regexp.MatchString("^web_main_live_", title)

	if match {

		substring := title[14:len(title)]

		fmt.Printf("Substring: %v",substring)

	} else {

		fmt.Println("Not found")
	}


}