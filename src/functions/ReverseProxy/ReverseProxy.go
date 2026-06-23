package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
)

func ReverseProxy(port string) {
	target, _ := url.Parse("http://127.0.0.1:5500")
	log.Println("Proxy target:", target.String())
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Host:", r.Host)
		proxy.ServeHTTP(w, r)
	})

	log.Println("Listening on :80")

	// log.Fatal(http.ListenAndServe(":80", nil))

	//port :443 for https and :80 for http
	server := &http.Server{
		Addr:    ":443",
		Handler: nil,
	}

	log.Println("Listening on :443")

	home := os.Getenv("HOME")
	if home == "" {
		var err error
		home, err = os.UserHomeDir()
		if err != nil {
			log.Fatal("could not resolve home directory:", err)
		}
	}
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser != "" {
		home = filepath.Join("/home", sudoUser)
	}
	certFile := filepath.Join(home, ".config/customcli/bin/promptshop.com.pem")
	keyFile := filepath.Join(home, ".config/customcli/bin/promptshop.com-key.pem")

	log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
}
