package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	listenAddr string
	serveDir   string
	port       string
)

func init() {
	flag.StringVar(&listenAddr, "l", "127.0.0.1", "Address to listen on")
	flag.StringVar(&serveDir, "d", "", "Directory to serve")
	flag.StringVar(&port, "p", "8000", "Port to listen on")
	flag.Parse()
}

func main() {
	log.Println("Starting server")
	log.Println("Running Configuration:")
	log.Printf(" - Listening on: %s:%s", listenAddr, port)
	log.Printf(" - Serving directory: %s", serveDir)

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(serveDir)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
