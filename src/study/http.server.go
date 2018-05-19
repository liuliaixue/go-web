package study

import (
	"log"
	"net/http"
)

func httpServer() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://www.baidu.com", 307)

	mux.Handle("/foo", rh)

	log.Println("Listenning...")

	http.ListenAndServe(":3000", mux)
}
