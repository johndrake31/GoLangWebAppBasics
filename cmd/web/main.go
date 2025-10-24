package main

import (
	"fmt"
	"net/http"

	"github.com/johndrake31/GoLangWebAppBasics/pkg/handlers"
)

const portNumber = ":8081"

// main is the main function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
