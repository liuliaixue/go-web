package study

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func httpServer() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://www.baidu.com", 307)

	mux.Handle("/foo", rh)

	log.Println("Listenning...")

	http.ListenAndServe(":3000", mux)

}

func httpServer2() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://www.baidu.com", 307)

	//handleDefault := http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("The time is: " )
	//})
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(res)
		res.Write([]byte("hello, go web\n"))
		res.Write([]byte("The time is: " + time.Now().Format(time.ANSIC)))
	})

	mux.Handle("/baidu", rh)

	log.Println("Listenning...")

	s := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
