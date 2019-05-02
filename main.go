package main
/*
reads - reads sql files
writes - output after modification
hot - files ready for import
*/
import (
	 "fmt"
	 "log"
	 "os"
	 cc "DbMigration/libraries"
)

var Why string = "migrate" 


var dir_reads string = "/home/hal/dumps/reads"
var	dir_writes string = "/home/hal/dumps/hot"
var tgtfilename string = "" 
var delete_write_bool bool = false
	

type data struct {
	infile string	
	outfile string
}

/*
func regexfunc(indata data) bool {
	//fmt.Printf("\n\tregexfunc:\n\tinfile: %v,\n\toutfile: %v\n",infile, outfile)

	if _, err := os.Stat(indata.infile); err == nil {
		//fmt.Printf("\ninfile: %v exists", infile)

		r, err := ioutil.ReadFile(indata.infile)//read file

		if err != nil {
			panic(err)
			return false
		}

		if _, err := os.Stat(indata.outfile); err == nil {//if outfile exists, deletes. 
			deleteFile(indata.outfile) 
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

//----
func deleteFile(infile string) bool {
    // deletes file, TEST3
    _, err := os.Stat(infile)

	if err == nil {

		err := os.Remove(infile)

		if err != nil {
			log.Printf("\tFile %v failed to delete.",infile)
			return false
		}
		//fmt.Println("\n\tFile Deleted")
		return true
	}

	return false
}

//----


//----
func newfileName(filename string) string {
	//find matches, if not return original name
	
	match, _ := regexp.MatchString("^web_main_live_", filename)

	if match {
		substring := filename[14:len(filename)]
		return substring
	}
	
	return filename
}
*/


//----
func main() {
	//TEST2
	var files []string 

	//LIVE
	//tgtfilename = "web_main_live_attorney_lead.sql" //single targeted sql file
	//delete_write_bool = true 

	//TODO todo 
	//if delete_write_bool {
	//	deleteFile(logfile)
	//}

	dir_reads = "/home/hal/dumps/hot"

	//logging
	loghandle, logerr := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)//appends to logfile


	if logerr != nil {
		panic(logerr)
	}

	log.SetOutput(loghandle)
	fmt.Printf("\n%v", "Starts processing sql files")
	fmt.Printf("\ndir: %v. %v\n", dir_reads, Why)
	log.Printf("%v", "\t------------\n\t\t\t\t\t\tStarts processing sql files")

	files = cc.WalkFiles(dir_reads)//returns file from directory
	result := cc.RegexVerifyfunc(&files)

	if result == false {
		fmt.Println("***Failed***")
	} else {
		fmt.Println("***Success!***")
	}

}


/*
func Verifyfunc(file_ptr *[]string) {
	
	var idx = 1
	fmt.Printf("\nfile size: %v\n", len(*file_ptr))
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists
			match, _ := regexp.MatchString("\\.sql$", infile)
			if match {
				fmt.Printf("%v. %v\n", idx, infile)
				cc.regexVerifyfunc(infile)
				idx += 1
			}
		}

	}//for
	
} 
*/



/*
	//fmt.Printf("\nfiles: %v size", len(files))
	var idx = 0
	for _, infile := range files {

		if _, err := os.Stat(infile); err == nil { //checks if file exists

			match, _ := regexp.MatchString("\\.sql$", infile)

			if match {//TEST2

				_, filename := path.Split(infile)

				if tgtfilename == "" || filename == tgtfilename {

					var dataobj data 
					dataobj.infile = infile
					dataobj.outfile = dir_writes + "/" + newfileName(filename)

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

					idx += 1
				}

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

	loghandle.Close() //close log handle
	*/

