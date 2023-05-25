package main

import (
	"fmt"
	"io"
	"net/http"
)

func HandleGetJoke(w http.ResponseWriter, r *http.Request) {
	const url = "https://geek-jokes.sameerkumar.website/api?format=json"
	resp, _ := http.Get(url)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	_, err := w.Write([]byte("take" + string(body) + "!\n"))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	handler := http.HandlerFunc(HandleGetJoke)

	mux.Handle("/joke", handler)

	fmt.Println("Start")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}

	fmt.Println("Finish")
}
