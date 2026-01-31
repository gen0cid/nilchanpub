package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx sync.Mutex

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

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Неорректный метод http-запроса")
		w.Write([]byte("Неорректный метод http-запроса"))
		return
	}

	httRequuestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Внутренняя ошибка сервера" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}

	httRequuestBodyString := string(httRequuestBody)

	mtx.Lock()
	d.names[d.nextId] = httRequuestBodyString
	d.nextId++
	mtx.Unlock()

	fmt.Printf("Имя: %s зарегистрировано на id: %d\n", httRequuestBodyString, d.nextId)
}

func (d *dataBase) handlerDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Неорректный метод http-запроса")
		w.Write([]byte("Неорректный метод http-запроса"))
		return
	}
	httRequuestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Внутренняя ошибка сервера" + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}

	httRequuestBodyString := string(httRequuestBody)

	httpRequuestId, err := strconv.Atoi(httRequuestBodyString)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		msg := "Некорректное число в теле запроса" + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	if _, ok := d.names[httpRequuestId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		msg := "элемент по указанному id не найден"
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	fmt.Printf("Удаление элемента c id:%d\n", httpRequuestId)
	mtx.Lock()
	delete(d.names, httpRequuestId)
	mtx.Unlock()
	fmt.Println("текущая мапа:", d.names)
}

func (d *dataBase) handlerListOneTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Неорректный метод http-запроса")
		w.Write([]byte("Неорректный метод http-запроса"))
		return
	}
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		msg := "Внутренняя ошибка сервера" + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	httpRequestBodyInt, err := strconv.Atoi(httpRequestBodyString)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		msg := "Некорректное число в теле запроса" + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	if _, ok := d.names[httpRequestBodyInt]; !ok {
		w.WriteHeader(http.StatusNotFound)
		msg := "элемент по указанному id не найден"
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	w.Write([]byte(httpRequestBodyString + " "))
	mtx.Lock()
	w.Write([]byte(d.names[httpRequestBodyInt]))
	mtx.Unlock()

}

func main() {
	db := newDataBase()

	http.HandleFunc("/Name", func(w http.ResponseWriter, r *http.Request) {
		db.handlerName(w, r)
	})

	http.HandleFunc("/Delete", func(w http.ResponseWriter, r *http.Request) {
		db.handlerDelete(w, r)
	})

	http.HandleFunc("/listOneName", func(w http.ResponseWriter, r *http.Request) {
		db.handlerListOneTask(w, r)
	})

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
