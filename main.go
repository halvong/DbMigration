package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"path"
	"compress/gzip"
	"regexp"
	"io/ioutil"
	"archive/tar"
	cc "DbMigration/libraries"
)

<<<<<<< HEAD
<<<<<<< HEAD
var version = "v1"
var which string
=======
var version = "v2"
>>>>>>> refs/heads/dev
var kind string = "qa"

var copy_dir string
<<<<<<< HEAD
var reads_dir string = "/home/hal/dumps/reads"
var writes_dir string = "/home/hal/dumps/hot"
var delete_dir string = "/home/hal/dumps/"
var copy_targz_dir string = "/home/hal/dumps/archives"
var destination_dir string = "/home/hal/Downloads/centos7work"
=======
=======
var version = "v2"
var kind string = "qa"

var copy_dir string
>>>>>>> dev
var reads_dir string = "/home/hal/dumps/reads/"
var hot_dir string = "/home/hal/dumps/hot/"
var delete_dir string = "/home/hal/dumps/"
var copy_targz_dir string = "/home/hal/dumps/archives/"
var targz_dir string = "/home/hal/dumps/" 
var dest_targz_dir = "/home/hal/dumps/archives/"
var destination_dir string = "/home/hal/Downloads/centos7work/"
var delete_infile_bool bool = false
<<<<<<< HEAD
>>>>>>> refs/heads/dev

<<<<<<< HEAD
var delete_infile_bool bool = false
	
=======
>>>>>>> refs/heads/dev
=======

>>>>>>> dev
func main() {
	current := time.Now()
	var files []string 
	var ok bool = false 

<<<<<<< HEAD
<<<<<<< HEAD
	fmt.Println("Starting")
	//0. copy from Dump folder to read folder  
	//which = "copy" 
	//copy_dir = "/home/hal/dumps/Dump20191107"
=======
=======
>>>>>>> dev
	fmt.Println("Starting script", current.Format("2006-01-02") )
	//var which string = ""
	//0. Copy from Dump folder to read folder  
	//var which = "copy" 
	//copy_dir = "/home/hal/dumps/Dump"+current.Format("20060102")
	//copy_dir = "/home/hal/dumps/Dump20191108"
<<<<<<< HEAD
>>>>>>> refs/heads/dev
=======
>>>>>>> dev

	//1. Migrate, process from read to hot folder
	//var which = "migrate"
	//kind = "qa" //qa or local
	//version = "v2"

	//2. Check
	//var which = "check" 
	//kind = "qa"//web_main_qa = qa; web_main_local = local

	//3. Grep
	//cd /home/hal/dumps/hot; grep -rni 'web_main_live' * 
	
	//4. Clean, deletes sql files only,  after upload is done
	//var which = "clean" //deletes all files in hot

<<<<<<< HEAD
<<<<<<< HEAD
	//5.
	//which = "delete"
	//delete_dir += "archives/Dump20191107"

	//6.
	//which = "copy_targz"
=======
=======
>>>>>>> dev
	//5. Delete Directory
	//var which = "delete"
	//delete_dir += "Dump"+current.Format("20060102")
	//delete_dir += "Dump20191115"
<<<<<<< HEAD
>>>>>>> refs/heads/dev

<<<<<<< HEAD
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
=======
=======

>>>>>>> dev
	//6. Copy Tar Gz
	//var which = "copy_targz"

	//7. Tar Gz
	//var which = "targz"
	//targz_dir += "Dump"+current.Format("20060102")
	//dest_targz_dir += "Dump"+current.Format("20060102")+".tar.gz" 
	//targz_dir += "Dump20191119" 
	//dest_targz_dir += "Dump20191119.tar.gz" 

	//checks for default folders/files
	if which == "copy" {
		ok = cc.CheckF([]string{"logs",copy_dir,reads_dir})
	} else if which == "migrate" || which == "clean" {
		ok = cc.CheckF([]string{"logs",reads_dir,hot_dir})
	} else if which == "delete" {
		ok = cc.CheckF([]string{"logs",delete_dir})
	} else if which == "copy_targz" {
		ok = cc.CheckF([]string{"logs",copy_targz_dir,destination_dir})
	} else if which == "targz" {
		ok = cc.CheckF([]string{"logs",targz_dir})
<<<<<<< HEAD
>>>>>>> refs/heads/dev
=======
>>>>>>> dev
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
<<<<<<< HEAD
<<<<<<< HEAD
			result = cc.RegexReadsfunc(&files, &delete_infile_bool, &writes_dir, &kind, &version)
=======
			result = cc.RegexReadsfunc(&files, &delete_infile_bool, &hot_dir, &kind, &version)
>>>>>>> refs/heads/dev
=======
			result = cc.RegexReadsfunc(&files, &delete_infile_bool, &hot_dir, &kind, &version)
>>>>>>> dev
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
<<<<<<< HEAD
<<<<<<< HEAD
			fmt.Println("\nDeletes all the files in hot.")	
			result = cc.DeleteFolder(&hot_dir)
=======
=======
>>>>>>> dev

			for _, folder := range([2]string{"hot_dir","reads_dir"}) {
				fmt.Printf("\nDeletes all the files in %v", folder)	
				result = cc.DeleteFolder(&folder)
			}

<<<<<<< HEAD
>>>>>>> refs/heads/dev
=======
>>>>>>> dev
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

<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> dev
	} else if which == "targz" {
		var fds []os.FileInfo

		cc.DeleteFile(&dest_targz_dir) //deletes logfile

		file, err := os.Create(dest_targz_dir)
		if err != nil { panic(err) }
		defer file.Close()
		// set up the gzip writer	
		gw := gzip.NewWriter(file)
		defer gw.Close()
		tw := tar.NewWriter(gw)
		defer tw.Close()

		if fds, err = ioutil.ReadDir(targz_dir); err != nil {
			panic(err)
		}
		var max int = len(fds) - 1

		for idx, fd := range fds {

			var srcfp string = path.Join(targz_dir, fd.Name())

			match, _ := regexp.MatchString(`\.sql$`, srcfp)

			if match {
				fmt.Printf("%v/%v. targz %v\n",idx, max, srcfp)

				if err := cc.AddFile(tw, srcfp); err != nil {
					panic(err)
				}
			}

		}//for
		
		result = true

<<<<<<< HEAD
>>>>>>> refs/heads/dev
=======
>>>>>>> dev
	} else {//check
		fmt.Println("Nothing is chosen")	
	}

	if result == false {
		fmt.Println("\n\n***Failed***")
	} else {
		fmt.Println("\n\n***Success!***")
	}

}


