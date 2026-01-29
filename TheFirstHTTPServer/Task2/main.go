package main

import (
	"fmt"
	"net/http"
)

func dogHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Гав"))
	if err != nil {
		fmt.Println("произошла ошибка:", err)
	}

	fmt.Println("dogHandler отработал успешно!")
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Мяу"))
	if err != nil {
		fmt.Println("произошла ошибка:", err)
	}

	fmt.Println("catHandler отработал успешно!")
}
func cowHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Муу"))
	if err != nil {
		fmt.Println("произошла ошибка:", err)
	}

	fmt.Println("cowHandler отработал успешно!")
}
func main() {
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/cat", catHandler)
	http.HandleFunc("/cow", cowHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("произошла ошибка:", err)
	}
	fmt.Println("программа успешно отработала!")
}
