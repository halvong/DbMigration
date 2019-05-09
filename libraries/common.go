package libraries

import (
	"fmt"
	"os"
	_ "io/ioutil"
	_ "regexp"
	"path/filepath"
)

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

func DeleteFile(infile_ptr *string) bool {
    // deletes file, TEST3
    _, err := os.Stat(*infile_ptr)

	if err == nil {

		err := os.Remove(*infile_ptr)

		if err != nil {
			fmt.Printf("\tFile %v failed to delete.", *infile_ptr)
			return false
		}
		return true
	}

	return false
}