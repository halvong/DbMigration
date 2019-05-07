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

	
func main() {
	//TEST2
	var which string = "migrate" 
	which = "check" 
	var files []string 
	var reads_dir string = "/home/hal/dumps/reads"
	var writes_dir string = "/home/hal/dumps/hot"
	var hot_dir string = writes_dir

	fmt.Println("Hello")

	//logging
	fp, logerr := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)//appends to logfile

	if logerr != nil {
		panic(logerr)
	}

	var delete_write_bool = true 

	//TODO todo 
	if delete_write_bool {
		cc.DeleteFile("logfile")
	}

	log.SetOutput(fp)
	fmt.Printf("\n%v", "Starts processing sql files")
	fmt.Printf("\ndir: %v. %v\n", reads_dir, which)
	
	log.Printf("%v", "\t------------\n\t\t\t\t\t\tStarts processing sql files")

	var result bool = false  
	if which != "migrate" {
		files = cc.WalkFiles(hot_dir)//returns file from directory
		result = cc.RegexVerifyHotfunc(&files)
	} else {
		files = cc.WalkFiles(reads_dir)//returns file from directory
		result = cc.RegexReadsfunc(&files, reads_dir, writes_dir)
	}

	if result == false {
		fmt.Println("***Failed***")
	} else {
		fmt.Println("***Success!***")
	}

}





