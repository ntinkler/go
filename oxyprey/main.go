/*
Oxyprey: MITM Proxy server written in GoLang designed to be integrated with the Ionic LLC platform.
Author: Nate Tinkler
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
  "log"
  "net"
  "net/http"
  "time"
  "github.com/davecgh/go-spew/spew"
  "fmt"
)

type timeHandler struct {
  	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  	tm := time.Now().Format(th.format)
  	w.Write([]byte("The time is: " + tm))
}

func createHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {


		//spew.Dump(r)
		if(validAddress(r.Host)) {
			_, err := net.Dial("tcp", r.Host)
			if(err != nil) {
				spew.Dump(err)
			} else {
				fmt.Println("Connected to " + r.Host)
			}
			
		}

		//spew.Dump(conn)

		//targetUrl := r.Host + r.URL.String()
		//spew.Dump(targetUrl)

    	w.Write([]byte("The time is: over..."))
  	}
  	return http.HandlerFunc(fn)
}

func validAddress(addr string) bool {
	fmt.Println(addr)
	host, port, err := net.SplitHostPort(addr)
	spew.Dump(err);
	fmt.Println(host + "~:~" + port)
	return true
}


func main() {
  	mux := http.NewServeMux()

  	rh := http.RedirectHandler("http://example.org", 307)
  	mux.Handle("/foo", rh)

	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/time", th)

	mux.Handle("/test", createHandler())

  	log.Println("Listening...")
  	http.ListenAndServe(":3000", createHandler())
}