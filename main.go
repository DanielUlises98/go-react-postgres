package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler finished")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(rw, "Hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(rw, err.Error(), internalError)
	}

}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}
