package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	log.Print("listening to - ", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to lister to server", err)
	}

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
