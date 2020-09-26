// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	flagHTTP = flag.String("http", ":8080", "HTTP port")
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*flagHTTP, nil))
}
