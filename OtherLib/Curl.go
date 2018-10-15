package OtherLib

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//go run main.go https://www.baidu.com/
func Curl01() {
	if len(os.Args) != 2 {
		fmt.Println("curl need two args")
		os.Exit(-1)
	}
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

}

//go run main.go https://www.baidu.com/
func Curl02() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
