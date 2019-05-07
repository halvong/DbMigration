package libraries

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"regexp"
	"path"
	_ "path/filepath"
)

var re = regexp.MustCompile("CREATE DATABASE  IF NOT EXISTS `web_main_live` \\/\\*!40100 DEFAULT CHARACTER SET latin1 \\*\\/;")
var re2 = regexp.MustCompile("web_main_live")

type data struct {
	infile string	
	outfile string
}

func RegexReadsfunc(file_ptr *[]string, reads_dir string, writes_dir string) bool {

	fmt.Printf("\nfiles: %v size", len(*file_ptr))
	var idx = 0
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists

			match, _ := regexp.MatchString("\\.sql$", infile)

			if match {//TEST2

				_, filename := path.Split(infile)


				var dataobj data 
				dataobj.infile = infile
				dataobj.outfile = writes_dir + "/" + newfileName(filename)

				/*	
				log.Printf("%v. Processing file: %v\n",strconv.Itoa(idx + 1),infile)
				fmt.Printf("\n%v. Processing file: %v",strconv.Itoa(idx + 1),infile)
				//var ok bool = regexfunc(infile, outfile)
				var ok bool = regexfunc(dataobj)

				if ok {

					log.Println("\tSuccess. Modify file:",dataobj.outfile)
					if delete_write_bool { 
						deleteFile(infile)
					}

				} else {
					log.Printf("\tFailed to modify %v",dataobj.outfile)
				}
				*/

				idx += 1

			}

		} else if os.IsNotExist(err) {
	        log.Printf("File %v not found.",infile)
		}

    }//for
    
    if idx == 0 {
		log.Println("No sql files found.\n\t\t\t\t\t\t-------")
		fmt.Println("\nNo sql files found.\n")
    } else {
		log.Println("Done\n\t\t\t\t\t\t-------")
		fmt.Println("\nDone\n")
    } 

	

	return true
}

func RegexVerifyHotfunc(file_ptr *[]string) bool {
	
	max := len(*file_ptr) 
	max -= 1

	var idx = 1
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists
			match, _ := regexp.MatchString("\\.sql$", infile)
			if match {

				fmt.Printf("%v/%v. %v\n", idx, max, infile)
				r, err := ioutil.ReadFile(infile)//read file

				if err != nil {
					panic(err)
					return false
				}

				result := re2.MatchString(string(r))
				if result == true {
					fmt.Printf("%v is wronged: %v.\n", infile, result)	
					return false
				} else {
					fmt.Printf("%v\n", result)	
				}	

				idx += 1
			}
		}

	}//for

	return true	
}

func newfileName(filename string) string {
	//find matches, if not return original name
	
	match, _ := regexp.MatchString("^web_main_live_", filename)

	if match {
		substring := filename[14:len(filename)]
		return substring
	}
	
	return filename
}