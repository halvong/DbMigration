package libraries

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"regexp"
	"strconv"
	"path"
)

var re_replaced = regexp.MustCompile("CREATE DATABASE  IF NOT EXISTS `web_main_live` \\/\\*!40100 DEFAULT CHARACTER SET latin1 \\*\\/;")
var re_live = regexp.MustCompile("web_main_live")
var re_qa = regexp.MustCompile("web_main_qa")
var re_local = regexp.MustCompile("local_web_main")
//var re_version = regexp.MustCompile("-- Server version	5.5.53-log")
var re_version = regexp.MustCompile(`-- Server version\s+.+`)

type data struct {
	infile string	
	outfile string
}

func RegexReadsfunc(file_ptr *[]string, delete_infile_ptr *bool, writes_dir_ptr *string, kind_ptr *string, version_ptr *string) bool {
	//1
	var max int = len(*file_ptr)
	max -= 1 //minus the folder
	fmt.Printf("\nfiles: %v size", max) 

	var idx = 0
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists

			match, _ := regexp.MatchString("\\.sql$", infile)

			if match {//TEST2

				_, filename := path.Split(infile)

				var dataobj data 
				dataobj.infile = infile
				dataobj.outfile = *writes_dir_ptr + "/" + newfileName(filename)

				log.Printf("%v/%v. Processing file: %v\n",strconv.Itoa(idx + 1), max, infile)
				fmt.Printf("\n%v/%v. Processing file: %v",strconv.Itoa(idx + 1), max, infile)
				
				var ok bool = regexfunc(dataobj, kind_ptr, version_ptr)//2

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

func RegexVerifyLivefunc(file_ptr *[]string) bool {
	
	max := len(*file_ptr) 
	max -= 1 //minus the folder

	var idx = 1
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists
			match, _ := regexp.MatchString("\\.sql$", infile)
			if match {

				r, err := ioutil.ReadFile(infile)//read file

				if err != nil {
					panic(err)
					return false
				}

				result := re_live.MatchString(string(r))
				if result == true {
					fmt.Printf("%v failed. Found web_main_live.\n", infile)	
					return false
				} 

				idx += 1
			}
		}

	}//for

	return true	
}

func RegexVerifyQALocalfunc(file_ptr *[]string, kind_ptr *string) bool {
	
	fmt.Println("")
	max := len(*file_ptr) 
	max -= 1 //minus the folder

	var idx = 1
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists
			match, _ := regexp.MatchString("\\.sql$", infile)
			if match {

				r, err := ioutil.ReadFile(infile)//read file

				if err != nil {
					panic(err)
					return false
				}

				if *kind_ptr == "local" {

					result := re_local.MatchString(string(r))
					if result == false {
						fmt.Printf("Not found local_web_main in %v\n", infile)	
						return false
					} 

				} else {//QA

					result := re_qa.MatchString(string(r))
					if result == false {
						fmt.Printf("Not found web_main_qa in %v\n", infile)	
						return false
					} 

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

func regexfunc(indata data, kind_ptr *string, version_ptr *string) bool {
	//2
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
		content := re_replaced.ReplaceAllString(string(r), "") 
		//2o2
		//re_live := regexp.MustCompile("web_main_live")
		if *kind_ptr == "local" { 
			content = re_live.ReplaceAllString(content, "local_web_main") 
		} else {

			content = re_live.ReplaceAllString(content, "web_main_qa") 

			if *version_ptr == "v2" { content = re_version.ReplaceAllString(content, "USE web_main_qa;") }
		}

		err = ioutil.WriteFile(indata.outfile, []byte(content), 0777)//writes content

		if err != nil {
			panic(err)
			return false
		}
		
		return true

	}//outside if 

	return false 
}
