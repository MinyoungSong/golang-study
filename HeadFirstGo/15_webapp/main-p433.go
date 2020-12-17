package main

import (
	"fmt"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {

	write(writer, "Hello, Web!")

}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {

	write(writer, "Salut, Web!")

}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {

	write(writer, "Namaste, Web!")

}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	fmt.Println("Listening !!!")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
