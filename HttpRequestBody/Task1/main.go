package main

import (
	"fmt"
	"io"
	"net/http"
)

var names []string

func handlerName(w http.ResponseWriter, r *http.Request) {
	httRequuestBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("неорректный ввод:", err)
	}
	httRequuestBodyString := string(httRequuestBody)
	names = append(names, httRequuestBodyString)

	fmt.Println("Ваше имя:", httRequuestBodyString)
	fmt.Println(names)

}

func main() {
	http.HandleFunc("/Name", handlerName)

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		fmt.Println("Ошибка:", err)
	}

}
