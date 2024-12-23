package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/urpc/uio/examples/httpsvr/uhttp"
)

func main() {
	var port int
	var loops int

	flag.IntVar(&port, "port", 8080, "server port")
	flag.IntVar(&loops, "loops", 0, "num loops")
	flag.Parse()

	var res = []byte("Hello World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(res)
	})

	err := uhttp.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if nil != err {
		log.Fatal(err)
	}
}
