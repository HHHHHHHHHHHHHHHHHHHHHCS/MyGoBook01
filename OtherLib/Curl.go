package OtherLib

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


//go run main.go https://www.baidu.com/
func Curl() {
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
