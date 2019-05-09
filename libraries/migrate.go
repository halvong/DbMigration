package libraries

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"regexp"
	"strconv"
	"path"
	//cc "common"
)

var re = regexp.MustCompile("CREATE DATABASE  IF NOT EXISTS `web_main_live` \\/\\*!40100 DEFAULT CHARACTER SET latin1 \\*\\/;")
var re2 = regexp.MustCompile("web_main_live")

type data struct {
	infile string	
	outfile string
}

func RegexReadsfunc(file_ptr *[]string, delete_infile_ptr *bool, writes_dir_ptr *string) bool {

	fmt.Printf("\nfiles: %v size", len(*file_ptr))
	var idx = 0
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists

			match, _ := regexp.MatchString("\\.sql$", infile)

			if match {//TEST2

				_, filename := path.Split(infile)

				var dataobj data 
				dataobj.infile = infile
				dataobj.outfile = *writes_dir_ptr + "/" + newfileName(filename)

				log.Printf("%v. Processing file: %v\n",strconv.Itoa(idx + 1),infile)
				fmt.Printf("\n%v. Processing file: %v",strconv.Itoa(idx + 1),infile)
				
				var ok bool = regexfunc(dataobj)

				if ok {

					log.Println("\tSuccess. Modify file:",dataobj.outfile)

					if *delete_infile_ptr == true { 
						DeleteFile(&infile) 
					}

				} else {
					log.Printf("\tFailed to modify %v",dataobj.outfile)
				}

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

func regexfunc(indata data) bool {
	//fmt.Printf("\n\tregexfunc:\n\tinfile: %v,\n\toutfile: %v\n",infile, outfile)

	if _, err := os.Stat(indata.infile); err == nil {
		//fmt.Printf("\ninfile: %v exists", infile)

		r, err := ioutil.ReadFile(indata.infile)//read file

		if err != nil {
			panic(err)
			return false
		}

		//if outfile exists, deletes. 
		if _, err := os.Stat(indata.outfile); err == nil {
			DeleteFile(&indata.outfile) 
		}		
		
		//regex string
		//content := strings.Replace(string(r), "CREATE DATABASE  IF NOT EXISTS `web_main_live`","", -1)
		//content = strings.Replace(content, "`web_main_live`","`web_main_qa`", -1)
		//re := regexp.MustCompile("CREATE DATABASE  IF NOT EXISTS `web_main_live`")

		//1o2
		//re := regexp.MustCompile("CREATE DATABASE  IF NOT EXISTS `web_main_live` \\/\\*!40100 DEFAULT CHARACTER SET latin1 \\*\\/;")
		content := re.ReplaceAllString(string(r), "") 
		//2o2
		//re2 := regexp.MustCompile("web_main_live")
		content = re2.ReplaceAllString(content, "web_main_qa") 

		err = ioutil.WriteFile(indata.outfile, []byte(content), 0777)//writes content

		if err != nil {
			panic(err)
			return false
		}
		
		return true

	}//outside if 

	return false 
}
