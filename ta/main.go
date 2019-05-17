package main

import (
    "bufio"
    "encoding/csv"
    _ "encoding/json"
    "fmt"
    "io"
    _ "log"
    "os"
)

func main() {
	
	csvFile, err := os.Open("ta/ta3.csv")
	//_ = csvFile

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

		fmt.Printf("%v. %v\n",idx, line)
		idx += 1
	}
	
}






