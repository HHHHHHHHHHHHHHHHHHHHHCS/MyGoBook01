package Web

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Main_HelloWorld(){
	//http://localhost:9090/tt?url_long=111&url_long=222&tt=22
	http.HandleFunc("/",sayHelloName_HelloWorld)
	err:=http.ListenAndServe(":9090",nil)
	if err !=nil {
		log.Fatalln("ListenAndServer:",err)
	}
}

func sayHelloName_HelloWorld(w http.ResponseWriter,r *http.Request){
	fmt.Println("_________________________")
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v :=range r.Form{
		fmt.Println("Key",k)
		fmt.Println("val:",strings.Join(v," "))
	}
	fmt.Fprint(w,"Hello astaxie!")
}