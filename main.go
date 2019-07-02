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
	"time"
	cc "DbMigration/libraries"
)

var which string = "migrate" //default
var kind_ptr string = "other"

var copy_dir string = "/home/hal/dumps/Dump20190702"
var reads_dir string = "/home/hal/dumps/reads"
var writes_dir string = "/home/hal/dumps/hot"

var delete_infile_bool bool = true
	
func main() {
	//1.
	//which = "copy" 

	//2.
	//which = "check" 
	//kind = "local"

	current := time.Now()

	var hot_dir string = writes_dir
	var files []string 

	var logfile = "logs/log_" + current.Format("2006-01-02")+".txt"
	cc.DeleteFile(&logfile) //deletes logfile

	//logging
	fp, logerr := os.OpenFile(logfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)//appends to logfile

	if logerr != nil {
		panic(logerr)
	}

	
	log.SetOutput(fp)
	log.Printf("\t------------\n\t\t\t\t\t\tStarts %v", which)

	var result bool = false  

	if which == "copy" {
		fmt.Printf("\n%v", "Starts copying folders.")
		fmt.Printf("\n\tCopies %v to %v.\n\n", copy_dir, reads_dir)

		//deletes reads folder
		result = cc.DeleteFolder(reads_dir)
	
		if result == false {
			panic("Cannot delete reads folder")
		}
		cc.Copyfiles(copy_dir, reads_dir)
	
	} else if which == "migrate" {

		fmt.Printf("\n%v", "Starts migrating to QA files.")
		fmt.Printf("\n\tMigrates dir: %v.\n\n", reads_dir)

		files = cc.WalkFiles(reads_dir)//returns file from directory
		if(len(files) > 0) {
			result = cc.RegexReadsfunc(&files, &delete_infile_bool, &writes_dir, &kind_ptr)
		} else {
			fmt.Println("\tNo file found")
		}

	} else {//check
		fmt.Printf("\n%v", "Starts checking sql files.")
		fmt.Printf("\n\tCheck hot dir: %v.\n\n", hot_dir)

		files = cc.WalkFiles(hot_dir)//returns file from directory
		if(len(files) > 0) {

			result = cc.RegexVerifyLivefunc(&files)

			if result != false {
				result = cc.RegexVerifyQALocalfunc(&files, &kind_ptr)
			}

		} else {
			fmt.Println("\tNo file found")
		}
	}

	if result == false {
		fmt.Println("***Failed***")
	} else {
		fmt.Println("***Success!***")
	}

}


