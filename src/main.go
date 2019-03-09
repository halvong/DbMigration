package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"io/ioutil"
	"strconv"
	"path/filepath"
)

func main() {

	var files []string
	//root := "/home/hal/dumps/Dump20190308"
	root := "/home/hal/dumps/imports"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })	
	
	if err != nil {
		panic(err)
	}

	idx := 1
	for _, file := range files {

		if _, err := os.Stat(file); err == nil {

			match, _ := regexp.MatchString("\\.sql$", file)

			if match { 
				fmt.Println(strconv.Itoa(idx) + ". " + file) 

				r, err := ioutil.ReadFile(file)

				if err != nil {
					panic(err)
	            }

				fmt.Println(r)

				content := strings.Replace(string(r), "CREATE DATABASE  IF NOT EXISTS `web_main_live`","", -1)

				err = ioutil.WriteFile(file, []byte(content), 0)
				if err != nil {
					panic(err)
				}

				idx += 1
			}

		} else if os.IsNotExist(err) {
	        fmt.Println("Not found:" + file)
		}

    }//for

}



