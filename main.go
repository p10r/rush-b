package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	server := Server{8080, http.NewServeMux()}
	server.mux.HandleFunc("POST /data", server.dataHandler())

	addr := fmt.Sprintf("localhost:%v", server.port)
	log.Printf("Starting Server on %v", addr)

	http.ListenAndServe(addr, server.mux)
}

type Server struct {
	port int
	mux  *http.ServeMux
}

func (server *Server) dataHandler() func(writer http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("ERROR: could not read body")
		}
		defer r.Body.Close()

		w.WriteHeader(200)

		if r.ContentLength == 0 {
			log.Println("Request body was empty")
			return
		}

		log.Println(string(body))
	}
}
