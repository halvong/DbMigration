package libraries

import (
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
	"path/filepath"
)

var re = regexp.MustCompile("CREATE DATABASE  IF NOT EXISTS `web_main_live` \\/\\*!40100 DEFAULT CHARACTER SET latin1 \\*\\/;")
var re2 = regexp.MustCompile("web_main_live")

type data struct {
	infile string	
	outfile string
}

func RegexVerifyfunc(file_ptr *[]string) bool {
	
	max := len(*file_ptr) 
	max -= 1

	var idx = 1
	for _, infile := range *file_ptr {

		if _, err := os.Stat(infile); err == nil { //checks if file exists
			match, _ := regexp.MatchString("\\.sql$", infile)
			if match {

				fmt.Printf("%v/%v. %v\n", idx, max, infile)
				r, err := ioutil.ReadFile(infile)//read file

				if err != nil {
					panic(err)
					return false
				}

				result := re2.MatchString(string(r))
				if result == true {
					fmt.Printf("%v is wronged: %v.\n", infile, result)	
					return false
				} else {
					fmt.Printf("%v\n", result)	
				}	

				idx += 1
			}
		}

	}//for

	return true	
}

	/*
	if _, err := os.Stat(infile); err == nil {
		//fmt.Printf("\ninfile: %v exists", infile)

		r, err := ioutil.ReadFile(infile)//read file

		if err != nil {
			panic(err)
			return false
		}

		result := re2.MatchString(string(r))
		fmt.Printf("%v\n", result)	

		return true

	}//outside if 

	return false 
	*/


func WalkFiles(dir_reads string) []string {

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