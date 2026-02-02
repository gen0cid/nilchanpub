package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx sync.Mutex

type dataBase struct {
	names  map[int]message
	nextId int
}

func newDataBase() *dataBase {
	return &dataBase{
		names:  make(map[int]message),
		nextId: 1,
	}
}

type message struct {
	Title       string `json:"title"`
	Index       int    `json:"index"`
	Description string `json:"description"`
	Isurgent    bool   `json:"isurgent"`
}

func (d *dataBase) handlerMessages(w http.ResponseWriter, r *http.Request) {

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

	var msg message

	json.Unmarshal(httRequuestBody, &msg)

	mtx.Lock()
	d.names[d.nextId] = msg
	fmt.Printf("Сообщение: %+v зарегистрировано на id: %d\n", msg, d.nextId)
	d.nextId++
	mtx.Unlock()

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

func (d *dataBase) handlerListMessages(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Неорректный метод http-запроса")
		w.Write([]byte("Неорректный метод http-запроса"))
		return
	}

	indexParam := r.URL.Query().Get("index")
	isurgentParam := r.URL.Query().Get("isurgent")

	if indexParam != "" || isurgentParam != "" {

		indexParamInt, err := strconv.Atoi(indexParam)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("err:", err.Error())
			w.Write([]byte("err:" + string(err.Error())))
			return
		}

		isurgentParamBool, err := strconv.ParseBool(isurgentParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("err:", err)
			w.Write([]byte("err: " + err.Error()))
			return
		}

		mtx.Lock()
		for k, v := range d.names {
			if v.Index == indexParamInt || v.Isurgent == isurgentParamBool {
				fmt.Println("id:", k, "--->>> message:", v)
				w.Write([]byte(fmt.Sprintf("id: %d ---> message: %+v", k, v)))
			}
		}
		mtx.Unlock()
	} else {

		mtx.Lock()
		b, err := json.Marshal(d.names)
		mtx.Unlock()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("err:", err)
			return
		}

		w.Write([]byte(b))

	}

}

func main() {
	db := newDataBase()

	http.HandleFunc("/Messages", func(w http.ResponseWriter, r *http.Request) {
		db.handlerMessages(w, r)
	})

	http.HandleFunc("/Delete", func(w http.ResponseWriter, r *http.Request) {
		db.handlerDelete(w, r)
	})

	http.HandleFunc("/ListMessages", func(w http.ResponseWriter, r *http.Request) {
		db.handlerListMessages(w, r)
	})

	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
