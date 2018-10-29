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

func login_repeatform(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method",r.Method)
	if r.Method=="GET"{
		curtime:=time.Now().Unix()
		h:=md5.New()
		io.WriteString(h,strconv.FormatInt(curtime,10))
		token:=fmt.Sprintf("%x",h.Sum(nil))

		t,_:=template.ParseFiles("Web/html/form.html")
		t.Execute(w,token)
	}else{
		r.ParseForm()
		token:=r.Form.Get("token")
		if token!=""{
			//在list 查重
		}else{
			//直接跳出 可能是假登录
		}
		fmt.Println("username length",len(r.Form["username"][0]))
		fmt.Println("username",template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password",template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w,[]byte(r.Form.Get("username")))
		fmt.Println(token)
	}
}

func Main_RepeatForm() {
	http.HandleFunc("/login", login_repeatform)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe", err)
	}
}
