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

var version = "v1"
var which string
var kind string = "qa"

var copy_dir string
var reads_dir string = "/home/hal/dumps/reads"
var writes_dir string = "/home/hal/dumps/hot"

var delete_infile_bool bool = false
	
func main() {

	fmt.Println("Starting")
	//0. copy from Dump folder to read folder  
	//which = "copy" 
	//copy_dir = "/home/hal/dumps/Dump20191107"

	//1. migrate, process from read to hot folder
	//which = "migrate"
	//kind = "qa" //qa or local
	//version = "v2"

	//2.
	which = "check" 
	kind = "qa"//web_main_qa = qa; web_main_local = local

	//3.
	//cd /home/hal/dumps/hot; grep -rni 'web_main_live' * 
	
	//4. after upload is done
	//which = "clean" //deletes all files in hot

	current := time.Now()

	var hot_dir string = writes_dir
	var files []string 

	//ok, findfile := cc.CheckF([]string{"logs",reads_dir,copy_dir,hot_dir})
	//Server version ok, findfile := cc.CheckF([]string{"logs",copy_dir,hot_dir})
	//if ok == false {
	//	panic(findfile+" does not exists\n")
	//}

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

		fmt.Println("Deletes all the files in hot.")	
		result = cc.DeleteFolder(hot_dir)

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
			result = cc.RegexReadsfunc(&files, &delete_infile_bool, &writes_dir, &kind, &version)
		} else {
			fmt.Println("\tNo file found")
		}

	} else if which == "check" {

		files = cc.WalkFiles(hot_dir)//returns file from directory
		var max int = len(files)

		fmt.Printf("Starts checking %d files for '%s'", max, kind)
		fmt.Printf("\n\tCheck hot dir: %v\n", hot_dir)

		if(len(files) > 0) {

			fmt.Printf("\tVerify 'no' Live")
			result = cc.RegexVerifyLivefunc(&files)

			if result != false {
				fmt.Printf("\n\tVerify '%s'", kind)
				result = cc.RegexVerifyQALocalfunc(&files, &kind)
			}

		} else {
			fmt.Println("\tNo file found")
		}

	} else if which == "clean" {

		fmt.Println("Deletes all the files in hot.")	
		result = cc.DeleteFolder(hot_dir)
	
		if result == false {
			panic("Cannot delete reads folder")
		}
	
	} else {//check
		fmt.Println("Nothing is chosen")	
	}

	if result == false {
		fmt.Println("***Failed***")
	} else {
		fmt.Println("***Success!***")
	}

}


