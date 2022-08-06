package main

import (
	"api-server/pkg/activo"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc('/', activo.Index)
	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
