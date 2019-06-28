package main 
//https://socketloop.com/tutorials/golang-archive-directory-with-tar-and-gzip

import (
	"os"
	"fmt"
	"flag"
)

func main() {

	flag.Parse() // get the arguments from command line

	destinationfile := flag.Arg(0)

	if destinationfile == "" {
		 fmt.Println("Usage : gotar destinationfile.tar.gz folder")
		 os.Exit(1)
	}

	sourcedir := flag.Arg(1)

	if sourcedir == "" {
		 fmt.Println("Usage : gotar destinationfile.tar.gz folder")
		 os.Exit(1)
	}

	fmt.Println("Hello World!")	
}

 func checkerror(err error) {

	 if err != nil {
		 fmt.Println(err)
		 os.Exit(1)
     }
 }