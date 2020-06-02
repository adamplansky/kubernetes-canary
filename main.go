package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	fmt.Fprintf(w, "version 1.0\n")
}

func main() {

	http.HandleFunc("/", rootHandler)
	fmt.Println("server is running :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println("fml")
	}
}
