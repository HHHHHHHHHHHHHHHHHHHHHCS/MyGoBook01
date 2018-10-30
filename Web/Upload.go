package Web

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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
		fmt.Println("Sprintf", fmt.Sprintf("%x", h.Sum(nil)))
		t, _ := template.ParseFiles("Web/html/upload.html")
		t.Execute(w, token)
	} else {
		token := r.Form.Get("token")
		if token != "" {
			//查重什么的

		} else {

		}
		r.ParseMultipartForm(1024 << 1) //byte
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprint(w, "%v", handler.Header)
		f, err := os.OpenFile("./Test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Println(token)
	}
}


//golang 模拟上传文件
func PostFile(filename, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func Main_Upload() {
	http.HandleFunc("/upload", login_upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe", err)
	}


}
func Main_UploadFile()  {
	target_url:="http://127.0.0.1:9090/upload"
	filename:="./Web/html/upload.html"
	PostFile(filename,target_url)
}



