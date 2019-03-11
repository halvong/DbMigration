package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"io/ioutil"
	"strconv"
	"path"
	"path/filepath"
)

func regexfunc(infile string, outfile string) bool {

	fmt.Printf("\nregexfunc:\n\tinfile: %v,\n\toutfile: %v\n",infile, outfile)

	if _, err := os.Stat(infile); err == nil {

		fmt.Printf("\ninfile: %v exists", infile)

		r, err := ioutil.ReadFile(infile)//read file

		if err != nil {
			panic(err)
			return false
		}

		//fmt.Println(r)

		//deletes
		content := strings.Replace(string(r), "CREATE DATABASE  IF NOT EXISTS `web_main_live`","", -1)
		content = strings.Replace(content, "`web_main_live`","`web_main_qa`", -1)

		err = ioutil.WriteFile(outfile, []byte(content), 0)//writes content

		if err != nil {
			panic(err)
			return false
		}
		
		return true
	} 

	return false 
}

func main() {

	var files []string
	dir_reads := "/home/hal/dumps/reads"
	dir_writes := "/home/hal/dumps/writes"

	err := filepath.Walk(dir_reads, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })	
	
	if err != nil {
		panic(err)
	}

	idx := 1
	for _, infile := range files {

		if _, err := os.Stat(infile); err == nil { //checks if file exists

			match, _ := regexp.MatchString("\\.sql$", infile)

			if match {//TODO

				filepath, filename := path.Split(infile)

				if filename == "web_main_live_attorney_lead.sql" {

					outfile := dir_writes + "/" + filename  

					fmt.Printf("%v. file_name: %v, path: %v\n",strconv.Itoa(idx), filename, filepath)
					var ok bool = regexfunc(infile, outfile)

					if ok {
						fmt.Println("\n\nSuccess. Writes file to",outfile)
					} else {
						fmt.Println("\nFailed")
					}

				}

				idx += 1
			}

		} else if os.IsNotExist(err) {
	        fmt.Println("Not found:" + infile)
		}

    }//for

}



