package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func numberHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "fail to read http body" + err.Error()

		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail to write http response", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	httpRequestBodyInt, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "fail to conv Atoi" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
	}

	if httpRequestBodyInt == 404 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if httpRequestBodyInt == 200 {
		w.WriteHeader(http.StatusOK)
		return
	}
	if httpRequestBodyInt == 201 {
		w.WriteHeader(http.StatusCreated)
		return
	}
	if httpRequestBodyInt == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if httpRequestBodyInt == 409 {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if httpRequestBodyInt == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/number", numberHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("ошибка:", err)
	}

}
