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

var which string = "migrate" 
var reads_dir string = "/home/hal/dumps/reads"
var writes_dir string = "/home/hal/dumps/hot"
	
func main() {
	//TEST2
	which = "check" 
	var hot_dir string = writes_dir
	var files []string 

	//logging
	fp, logerr := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)//appends to logfile

	if logerr != nil {
		panic(logerr)
	}

	var delete_infile_bool bool = true

	//deletes logfile
	//cc.DeleteFile("logfile")
	
	log.SetOutput(fp)
	fmt.Printf("\n%v", "Starts processing sql files")
	fmt.Printf("\n\tRead dir: %v. Type: %v\n", reads_dir, which)
	
	log.Printf("%v", "\t------------\n\t\t\t\t\t\tStarts processing sql files")

	var result bool = false  

	if which == "migrate" {
		files = cc.WalkFiles(reads_dir)//returns file from directory
		result = cc.RegexReadsfunc(&files, &delete_infile_bool, &writes_dir)
	} else {
		files = cc.WalkFiles(hot_dir)//returns file from directory
		result = cc.RegexVerifyHotfunc(&files)
	}

	if result == false {
		fmt.Println("***Failed***")
	} else {
		fmt.Println("***Success!***")
	}

}





