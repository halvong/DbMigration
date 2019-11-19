package libraries

import (
	"os"
	"io"
	"fmt"
	"path"
	"regexp"
	"io/ioutil"
	"path/filepath"
	"archive/tar"
	
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

func CopyTargzfiles(src_ptr *string, dest_ptr *string) bool {
	
	var ok bool = true
	var err error
	var fds []os.FileInfo
	var src_arr []string 
	var dest_arr []string 
	var max int 

	if fds, err = ioutil.ReadDir(*src_ptr); err != nil {
		panic(err)
		return false
	}
	
	
	for _, fd := range fds {
		var srcfp string = path.Join(*src_ptr, fd.Name())
		var dstfp string = path.Join(*dest_ptr, fd.Name())

		match, _ := regexp.MatchString(`\.tar\.gz$`, srcfp)

		if match {
			src_arr = append(src_arr, srcfp)		
			dest_arr = append(dest_arr, dstfp)		
		}
	}//for

	max = len(src_arr)	
	for idx, ffile := range(src_arr) {

		fmt.Printf("%v/%v. Copying %v => %v\n",idx+1, max, ffile, dest_arr[idx])

		DeleteFile(&dest_arr[idx])
		if err = File(ffile, dest_arr[idx]); err != nil {
			ok = false
			panic(err)	
		}	
	}

	if ok {
		//deletes file	
		for _, ffile := range(src_arr) {
			DeleteFile(&ffile)
		}
	}

	return ok
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

func DeleteFolder(dir_read_ptr *string) bool {
	var ok bool = true
	var files []string = WalkFiles(*dir_read_ptr)

	for _, file := range files {

		if len(file) > 0 {
			ok := DeleteFile(&file)
			if ok == false { break }
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
	}

	return true
}

func CheckF(arr []string)bool {

	var ok bool = true  

	for _, infile := range(arr) {	

		//fmt.Println(idx,infile)
		 _, err := os.Stat(infile)

		if err != nil { 
			fmt.Println("Line134",err)
			ok = false
			break		
		}

		if os.IsNotExist(err) { 
			fmt.Println("2.")
			ok = false
			break		
		}
	}

	return ok
}

func RemoveDirectory(folder_ptr *string) bool {

    _, err := os.Stat(*folder_ptr)

	if err == nil {

		var ok bool = DeleteFolder(folder_ptr)

		if ok {

			err := os.Remove(*folder_ptr)

			if err != nil {
				fmt.Printf("\tFolder %v failed to delete.", *folder_ptr)
				fmt.Printf("\tError: %v.", err)
				return false
			}

		} else {
			return false
		}
	}

	return true
}
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> dev

func AddFile(tw * tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball	
		header := new(tar.Header)
		header.Name = path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball 
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}

<<<<<<< HEAD
>>>>>>> refs/heads/dev
=======
>>>>>>> dev
