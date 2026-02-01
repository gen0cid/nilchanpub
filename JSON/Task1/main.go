package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name    string  `json:"name"`
	Adress  string  `json:"adress"`
	Age     int     `json:"age"`
	Married bool    `json:"married"`
	Height  float64 `json:"height"`
}

func personInHandler(w http.ResponseWriter, r *http.Request) {
	httoRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var person Person
	if err := json.Unmarshal(httoRequestBody, &person); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(person)
}

func personOutHandler(w http.ResponseWriter, r *http.Request) {
	pers := Person{
		Name:    "Геннадий",
		Adress:  "проспект ленина",
		Age:     22,
		Married: false,
		Height:  177.0,
	}

	b, err := json.Marshal(pers)

	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/PersonIn", personInHandler)
	http.HandleFunc("/PersonOut", personOutHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
}
