DbMigration, Golang, working ubuntu
11/14, Thurs

learning gzip & tar
misc.go

#path
cd /home/hal/Documents/softwares/go-eclipse/workspace/src/DbMigration

#checks "no" live
cd /home/hal/dumps/hot; grep -rni 'web_main_live' * 


#datatypes
bool
string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

#notes
fmt.Println("\ntype: ", reflect.TypeOf(fp))	