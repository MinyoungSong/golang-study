package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {

	message := []byte("Hello, Web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/hello", viewHandler)
	fmt.Println("Listening !!!")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
