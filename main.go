package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("port", ":8080", "host:port to listen on")
	dir  = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("Server: %v on %v\n", *dir, *port)
	log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir(*dir))))
}
