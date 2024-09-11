package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VincentDevi/JobsTrack/internal/view" // Correct import path
	"github.com/a-h/templ"
)

func main() {

	http.Handle("/", templ.Handler(view.Hello()))

	fmt.Println("Server running on port 8080 ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
