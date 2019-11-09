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
var delete_dir string = "/home/hal/dumps/"
var copy_targz_dir string = "/home/hal/dumps/archives"
var destination_dir string = "/home/hal/Downloads/centos7work"

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
	//which = "check" 
	//kind = "qa"//web_main_qa = qa; web_main_local = local

	//3.
	//cd /home/hal/dumps/hot; grep -rni 'web_main_live' * 
	
	//4. after upload is done
	//which = "clean" //deletes all files in hot

	//5.
	//which = "delete"
	//delete_dir += "archives/Dump20191107"

	//6.
	which = "copy_targz"

	fmt.Printf("%v", delete_dir)

	var hot_dir string = writes_dir
	var files []string 
	current := time.Now()

	//checks for default folders/files
	var ok bool = false 
	if which == "copy" {
		ok = cc.CheckF([]string{"logs",copy_dir,reads_dir})
	} else if which == "migrate" {
		ok = cc.CheckF([]string{"logs",reads_dir,hot_dir})
	} else if which == "clean" {
		ok = cc.CheckF([]string{"logs",hot_dir})
	} else if which == "delete" {
		ok = cc.CheckF([]string{"logs",delete_dir})
	} else if which == "copy_targz" {
		ok = cc.CheckF([]string{"logs",copy_targz_dir,destination_dir})
	} else {
		ok = cc.CheckF([]string{"logs"})
	}
	if ok == false {
		panic("Default file/folder does not exists\n")
	}

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
		result = cc.DeleteFolder(&hot_dir)

		fmt.Printf("\n%v", "Starts copying folders.")
		fmt.Printf("\n\tCopies %v to %v.\n\n", copy_dir, reads_dir)

		//deletes reads folder
		result = cc.DeleteFolder(&reads_dir)
	
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
				fmt.Printf("\n\tVerify all '%s'", kind)
				result = cc.RegexVerifyQALocalfunc(&files, &kind)
			}

		} else {
			fmt.Println("\tNo file found")
		}

	} else if which == "clean" || which == "delete" {

		if which == "clean" {
			fmt.Println("\nDeletes all the files in hot.")	
			result = cc.DeleteFolder(&hot_dir)
		} else if which == "delete" {
			fmt.Println("\nDeletes "+delete_dir)	
			result = cc.RemoveDirectory(&delete_dir)
		}
	
		if result == false {
			panic("\nCannot delete reads folder")
		}
	
	} else if which == "copy_targz" {

		fmt.Printf("\n%v", "Starts copying folders.")
		fmt.Printf("\n\tCopies %v to %v.\n\n", copy_targz_dir, destination_dir)

		result = cc.CopyTargzfiles(&copy_targz_dir, &destination_dir)

	} else {//check
		fmt.Println("Nothing is chosen")	
	}

	if result == false {
		fmt.Println("***Failed***")
	} else {
		fmt.Println("***Success!***")
	}

}


