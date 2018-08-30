package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/login", loginFunc)
	err := http.ListenAndServe(":8080", nil)
	if err!=nil {
		fmt.Println("http listen failed")
	}
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello basic web"))
}

func loginFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("login"))
}
