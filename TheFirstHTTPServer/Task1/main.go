package main

import (
	"fmt"
	"net/http"
)

func handlerWelcome(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Здравствуйте!"))
	if err != nil {
		fmt.Println("произошла ошибка:", err)
	}
	fmt.Println("handlerWelcome отработал успешно!")
}
func main() {
	http.HandleFunc("/", handlerWelcome)

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		fmt.Println("Произошла ошибка:", err)
	}

}
