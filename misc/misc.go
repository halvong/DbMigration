package main

import (
	"os"
	"archive/tar"
	"fmt"
	"regexp"
	"path"	
	"io/ioutil"
	"io"
	"compress/gzip"
)

func addFile(tw * tar.Writer, path string) error {
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

func main() {
	// set up the output file
	var err error
	var fds []os.FileInfo
	file, err := os.Create("Dump.tar.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// set up the gzip writer	
	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	var targz_dir string = "/home/hal/dumps/archives/Dump20191115" 
	var src_arr []string


	if fds, err = ioutil.ReadDir(targz_dir); err != nil {
		panic(err)
	}
	var max int = len(fds) - 1

	for idx, fd := range fds {

		var srcfp string = path.Join(targz_dir, fd.Name())

		match, _ := regexp.MatchString(`\.sql$`, srcfp)

		if match {
			src_arr = append(src_arr, srcfp)
			fmt.Printf("%v/%v. targz %v\n",idx, max, srcfp)

			if err := addFile(tw, srcfp); err != nil {
				panic(err)
			}
		}

	}

	// add each file as needed into the current tar archive
	//for i := range paths {
		//fmt.Println(paths[i])
		//if err := addFile(tw, paths[i]); err != nil {
		//	log.Fatalln(err)
		//}
	//}
}