package main

import (
	"bookshelf/http"
	"bookshelf/methods"
	"fmt"
)

func main() {
	bookshelf := methods.NewBookShelf()
	httpHandlers := http.NewhttpHandlers(bookshelf)
	server := http.NewServer(httpHandlers)
	if err := server.Startserver(); err != nil {
		fmt.Println("failed to start http server")
		return
	}
}
