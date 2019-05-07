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


//var dir_reads string = "/home/hal/dumps/reads"
//var dir_writes string = "/home/hal/dumps/hot"
//var tgtfilename string = "" 
var delete_write_bool bool = false
	
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


*/


//----
func main() {
	//TEST2
	var files []string 
	var reads_dir string = "/home/hal/dumps/reads"
	var writes_dir string = "/home/hal/dumps/hot"

	fmt.Println("Hello")

	//logging
	fp, logerr := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)//appends to logfile

	if logerr != nil {
		panic(logerr)
	}

	delete_write_bool = false 

	//TODO todo 
	if delete_write_bool {
		cc.DeleteFile("logfile")
	}


	log.SetOutput(fp)
	fmt.Printf("\n%v", "Starts processing sql files")
	fmt.Printf("\ndir: %v. %v\n", reads_dir, Why)
	
	log.Printf("%v", "\t------------\n\t\t\t\t\t\tStarts processing sql files")

	files = cc.WalkFiles(reads_dir)//returns file from directory
	_ = files
	//result := cc.RegexVerifyHotfunc(&files)
	result := cc.RegexReadsfunc(&files, reads_dir, writes_dir)

	//var result bool = false  
	if result == false {
		fmt.Println("***Failed***")
	} else {
		fmt.Println("***Success!***")
	}

	//log.Close() //close log handle
}





