package Web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayHelloName_Form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hello astaxie!")
}

func login_form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Web/html/form.html")
		t.Execute(w, nil)
	} else {
		safe(w,r)
	}
}

func unsafe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("username", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
	fmt.Fprint(w, r.Form["username"], r.Form["password"])
}

func safe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("username", template.HTMLEscapeString(r.Form.Get("username")))
	fmt.Println("password:",template.HTMLEscapeString(r.Form.Get("password")))
	template.HTMLEscape(w,[]byte(r.Form.Get("username")))

	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", template.HTML("<script>alert('you have been pwned')</script>"))
	fmt.Println()
	fmt.Println(t,err)
}

func Main_Form() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login_form)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)

	}
}
