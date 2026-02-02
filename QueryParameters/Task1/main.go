package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fooPar := r.URL.Query().Get("foo")
	booPar := r.URL.Query().Get("boo")

	fmt.Println("foopar:", fooPar)
	fmt.Println("boopar:", booPar)
	w.Write([]byte("foopar: " + fooPar + "\n"))
	w.Write([]byte("boopar: " + booPar + "\n"))

}
func main() {
	http.HandleFunc("/def", handler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("err:", err)
		return
	}
}
