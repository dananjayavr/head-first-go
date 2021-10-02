package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request)  {
	write(writer, "Hello!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request)  {
	write(writer,"Salut!")
}

func germanHandler(writer http.ResponseWriter, request *http.Request)  {
	write(writer, "Halo!")
}

func viewHandler(writer http.ResponseWriter, request *http.Request)  {
	message := []byte("Hello, Web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/en/hello", englishHandler)
	http.HandleFunc("/fr/hello", frenchHandler)
	http.HandleFunc("/de/hello", germanHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
