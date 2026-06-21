package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(port string) {
	target, _ := url.Parse("http://127.0.0.1:5173")
	log.Println("Proxy target:", target.String())
	fmt.Println("now doing revese proxy")
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Host:", r.Host)
		proxy.ServeHTTP(w, r)
	})

	log.Println("Listening on :80")

	log.Fatal(http.ListenAndServe(":80", nil))
}
