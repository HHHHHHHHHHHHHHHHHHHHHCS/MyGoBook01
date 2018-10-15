package main

import (
	"os"
)

var(
	Stdin = newfile
)

func main()  {

}

func NewFile(fd uintptr,name string) *os.File{
	fdi :=int(fd)
	if fdi<0{
		return  nil
	}
	f:=&File{&file{}}
}