package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"study"
	"time"
	// "study"
)

func main() {
	fmt.Println("main is running")
	server()
	// study.Entry()
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的

}

func sayhelloDate(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, time.Now().Format("2006-01-02 15:04:05")) //这个写入到w的是输出到客户端的

}

func addNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.Body)
	if r.Method == "POST" {
		// r.ParseForm()
		fmt.Println(r.PostForm)
		fmt.Println(r.PostForm == nil)
		var (
			userID  = r.PostFormValue("userID")
			content = r.PostFormValue("content")
		)
		fmt.Printf("userId is  : %s\n", userID)
		fmt.Printf("content is : %s\n", content)
		study.AddNote(userID, content)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("The time is: "))
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	res := study.GetNote()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Write([]byte("The time is: "))
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Marshal error")
	}

	w.Header().Set("content-type", "application/json")

	w.Write([]byte(b))
}

func server() {
	// http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/time", sayhelloDate) //设置访问的路由
	http.HandleFunc("/note/add", addNote)
	http.HandleFunc("/notes", getNotes)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
