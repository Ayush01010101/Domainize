package functions

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Ayush01010101/Custom-Domain-CLI.git/src/utlities"
)

func ReverseProxy(port string, domain string) {

	target, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%s", port))
	log.Println("Proxy target:", target.String())
	proxy := httputil.NewSingleHostReverseProxy(target)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Host:", r.Host)
		proxy.ServeHTTP(w, r)
	})

	log.Println("Listening on :80")
	go func() {
		log.Fatal(http.ListenAndServe(":80", mux))
	}()

	//port :443 for https and :80 for http
	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
	}

	log.Println("Listening on :443")

	certFile, keyFile, err := utlities.GetSSL_TLS_keys(domain)
	if err != nil {
		log.Fatal(err)
	}

	server.ListenAndServeTLS(certFile, keyFile)
}
