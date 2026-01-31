package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/k0kubun/pp"
)

func handler(w http.ResponseWriter, r *http.Request) {

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "server error " + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	if _, ok := r.Header[httpRequestBodyString]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("not found el in map ")
		return
	}

	msg := "Привет!!!!! " + r.Header[httpRequestBodyString][0]
	pp.Printf(msg)

}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println(err)
		return
	}

}
