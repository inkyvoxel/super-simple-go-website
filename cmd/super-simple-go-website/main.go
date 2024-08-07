package main

import (
	"fmt"
	"net/http"

	"github.com/inkyvoxel/super-simple-go-website/web/app"
)

func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.IndexHandler)

	fmt.Println("Starting server at http://localhost:8080/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
