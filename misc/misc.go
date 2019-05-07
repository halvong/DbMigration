package main

import (
	"os"
	"fmt"
	"reflect"
)

func main() {
	fp, _ := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)//appends to logfile
	fmt.Println("\ntype: ", reflect.TypeOf(fp))
}