package Web

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func login_upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		fmt.Println("curtime", strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("Sprintf",fmt.Sprintf("%x", h.Sum(nil)))
		t, _ := template.ParseFiles("Web/html/upload.html")
		t.Execute(w, token)
	} else {
		token := r.Form.Get("token")
		if token != "" {
			//查重什么的
			r.ParseMultipartForm(1024<<1)//byte
			file,handler,err :=r.FormFile("uploadfile")
			if err !=nil{
				fmt.Println(err)
				return
			}

		} else {

		}

		fmt.Println(token)
	}
}

func Main_Upload() {
	http.HandleFunc("/upload", login_upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe", err)
	}
}
