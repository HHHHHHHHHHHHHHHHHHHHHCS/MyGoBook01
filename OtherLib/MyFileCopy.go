package OtherLib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputAddr  = `input`
	outputAddr = `output`
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}

func Main_FileCopy() {
	isExist, err := PathExists(outputAddr)
	if err != nil {
		panic(err)
	}
	if !isExist {
		os.Mkdir(outputAddr, os.ModePerm)
	}

	files, _ := ioutil.ReadDir(inputAddr)
	for _, f := range files {
		if strings.Contains(f.Name(), "(") {
			src, err := os.Open(inputAddr + "/" + f.Name())
			if err != nil {
				panic(err)
			}
			defer src.Close()

			//dest, err := os.OpenFile(outputAddr+"/"+f.Name(), os.O_CREATE|os.O_WRONLY, os.ModePerm)
			dest, err := os.Create(outputAddr + "/" + f.Name())
			if err != nil {
				panic(err)
			}
			defer dest.Close()

			nBytes, err := io.Copy(dest, src)

			fmt.Println(nBytes)



		}
	}
}
