package main
/*
reads - reads sql files
writes - output after modification
imports - files ready for import
*/
import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"regexp"
	"strconv"
	"path"
	"path/filepath"
)

func regexfunc(infile string, outfile string) bool {
	//fmt.Printf("\n\tregexfunc:\n\tinfile: %v,\n\toutfile: %v\n",infile, outfile)

	if _, err := os.Stat(infile); err == nil {
		//fmt.Printf("\ninfile: %v exists", infile)

		r, err := ioutil.ReadFile(infile)//read file

		if err != nil {
			panic(err)
			return false
		}

		if _, err := os.Stat(outfile); err == nil {//if outfile exists, deletes. 
			deleteFile(outfile) 
		}		
		
		//regex string
		//content := strings.Replace(string(r), "CREATE DATABASE  IF NOT EXISTS `web_main_live`","", -1)
		//content = strings.Replace(content, "`web_main_live`","`web_main_qa`", -1)

		re := regexp.MustCompile("CREATE DATABASE  IF NOT EXISTS `web_main_live`")
		content := re.ReplaceAllString(string(r), "") 

		re2 := regexp.MustCompile("`web_main_live`")
		content = re2.ReplaceAllString(content, "`web_main_qa`") 

		err = ioutil.WriteFile(outfile, []byte(content), 0777)//writes content

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
    var err = os.Remove(infile)

    if err != nil {
	    log.Printf("\tFile %v failed to delete.",infile)
        return false
    }
    //fmt.Println("\n\tFile Deleted")
	return true
}

//----
func walkFiles(dir_reads string) []string {

	var files []string

	err := filepath.Walk(dir_reads, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })	
	
	if err != nil {
		panic(err)
	}

	return files	
}

//--------------------------
func main() {
	//TEST1
	dir_reads := "/home/hal/dumps/reads"
	dir_writes := "/home/hal/dumps/writes"
	var tgtfilename = "" 
	//tgtfilename = "web_main_live_attorney_lead.sql"
	var delete_write_bool bool = false
	//delete_write_bool = true 

	var files []string

	loghandle, logerr := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if logerr != nil {
		panic(logerr)
	}
	log.SetOutput(loghandle)
	fmt.Printf("\n%v", "Starts processing sql files")
	log.Printf("%v", "\t------------\n\t\t\t\t\t\tStarts processing sql files")

	files = walkFiles(dir_reads)//returns file from directory

	//fmt.Printf("\nfiles: %v size", len(files))
	var idx = 0
	for _, infile := range files {

		if _, err := os.Stat(infile); err == nil { //checks if file exists

			match, _ := regexp.MatchString("\\.sql$", infile)

			if match {//TEST2

				_, filename := path.Split(infile)

				if tgtfilename == "" || filename == tgtfilename {

					outfile := dir_writes + "/" + filename  

					log.Printf("%v. file_name: %v\n",strconv.Itoa(idx + 1),infile)
					fmt.Printf("%v. file_name: %v\n",strconv.Itoa(idx + 1),infile)
					var ok bool = regexfunc(infile, outfile)

					if ok {

						log.Println("\tSuccess. Modify file:",outfile)
						if delete_write_bool { 
							deleteFile(infile)
						}

					} else {
						log.Printf("\tFailed to modify %v",outfile)
					}

					idx += 1
				}

			}

		} else if os.IsNotExist(err) {
	        log.Printf("File %v not found.",infile)
		}

    }//for
    
    if idx == 0 {
		msg := "\nNo sql files found."
		log.Println(msg)
		fmt.Println(msg)
    } else {
		msg := "Done\n"
		log.Println(msg)
		fmt.Println(msg)
    } 

	loghandle.Close() //close log handle
}
