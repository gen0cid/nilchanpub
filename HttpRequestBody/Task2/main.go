package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type dataBase struct {
	names  map[int]string
	nextId int
}

func newDataBase() *dataBase {
	return &dataBase{
		names:  make(map[int]string),
		nextId: 1,
	}
}

func (d *dataBase) handlerName(w http.ResponseWriter, r *http.Request) {
	httRequuestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("неорректный ввод:", err)
		return
	}

	httRequuestBodyString := string(httRequuestBody)

	d.names[d.nextId] = httRequuestBodyString
	fmt.Printf("Имя: %s зарегистрировано на id: %d\n", httRequuestBodyString, d.nextId)
	d.nextId++

}

func (d *dataBase) handlerDelete(w http.ResponseWriter, r *http.Request) {
	httRequuestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("неорректный ввод:", err)
		return
	}

	httRequuestBodyString := string(httRequuestBody)

	httpRequuestId, err := strconv.Atoi(httRequuestBodyString)
	if err != nil {
		fmt.Println("неорректный ввод, введите id:", err)
		return
	}
	_, ok := d.names[httpRequuestId]
	if !ok {
		fmt.Println("Нет такого элемента в мампе")
		return
	}

	fmt.Printf("Удаление элемента c id:%d\n", httpRequuestId)
	delete(d.names, httpRequuestId)
	fmt.Println("текущая мапа:", d.names)
}

func main() {
	db := newDataBase()

	http.HandleFunc("/Name", func(w http.ResponseWriter, r *http.Request) {
		db.handlerName(w, r)
	})

	http.HandleFunc("/Delete", func(w http.ResponseWriter, r *http.Request) {
		db.handlerDelete(w, r)
	})

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
