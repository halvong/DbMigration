package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"path/filepath"
)

func main() {

	var files []string
	root := "/home/hal/dumps/Dump20190308"

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
				idx += 1
			}

		} else if os.IsNotExist(err) {
	        fmt.Println("Not found:" + file)
		}

    }//for

}



