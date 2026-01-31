package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money int
var mtx sync.Mutex

func PayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	httpRequestBodyInt, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("incorrect input:", err)
		w.Write([]byte(err.Error()))
		return
	}
	mtx.Lock()
	if httpRequestBodyInt > money {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("У вас нет прав на выполнение этой операции (недостаточно средств)")
		fmt.Println("Ваш счет:", money)
		return
	}

	money -= httpRequestBodyInt

	fmt.Println("средства успешно сняты со счета!")
	fmt.Println("Сумма:", httpRequestBodyInt)
	fmt.Println("Ваш счет:", money)

	mtx.Unlock()
}

func main() {
	money = 1000

	mtx.Lock()
	fmt.Println("Ваш счет:", money)
	mtx.Unlock()

	http.HandleFunc("/pay", PayHandler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("error: ", err)
		return
	}
}
