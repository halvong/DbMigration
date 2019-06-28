package libraries

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"regexp"
	"path/filepath"
	"path"
)

func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	//reads source file
	if srcfd, err = os.Open(src); err != nil {
		panic(err)
	}
	defer srcfd.Close()

	//creates dest file
	if dstfd, err = os.Create(dst); err != nil {
		panic(err)
	}
	defer dstfd.Close()

	//copies
	if _, err = io.Copy(dstfd, srcfd); err != nil {
		panic(err)
	}

	//get source directory stats
	if srcinfo, err = os.Stat(src); err != nil {
		panic(err)
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func Copyfiles(src string, dest string) error {
	var err error
	var fds []os.FileInfo

	if fds, err = ioutil.ReadDir(src); err != nil {
		panic(err)
	}
	var max int = len(fds) - 1

	for idx, fd := range fds {

		var srcfp string = path.Join(src, fd.Name())
		var dstfp string = path.Join(dest, fd.Name())

		fmt.Printf("%v/%v. Copying %v => %v\n",idx, max, srcfp, dstfp)

		if err = File(srcfp, dstfp); err != nil {
			panic(err)
		}	

	}//for

	return nil
}

func WalkFiles(dir_reads string) []string {

	var files []string

	err := filepath.Walk(dir_reads, func(path string, info os.FileInfo, err error) error {

		match, _ := regexp.MatchString("\\.sql$", path)

		if match {
	        files = append(files, path)
		}

        return nil
    })	

	if err != nil {
		panic(err)
	}

	return files	
}

func DeleteFolder(dir_read string) bool {
	var ok bool = true
	var files []string = WalkFiles(dir_read)

	for _, file := range files {
		ok := DeleteFile(&file)

		if ok == false {
			break	
		}
	}

	return ok
}

func DeleteFile(infile_ptr *string) bool {
    // deletes file
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