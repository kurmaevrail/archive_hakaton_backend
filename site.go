package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	if write() == "" {
		fmt.Fprintf(w, "Bad")
	} else {
		fmt.Fprintf(w, write())
	}
}

func site() {
	http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера

	// Если хотите использовать https, то вместо ListenAndServe используйте ListenAndServeTLS
	// err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
