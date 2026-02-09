package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	httpHandlers *httpHandlers
}

func NewServer(httpHandlers *httpHandlers) *httpServer {
	return &httpServer{
		httpHandlers: httpHandlers,
	}
}
func (s *httpServer) Startserver() error {
	router := mux.NewRouter()

	router.Path("/books").Methods(http.MethodPost).HandlerFunc(s.httpHandlers.AddBookHandler)
	router.Path("/books/{title}").Methods(http.MethodPatch).HandlerFunc(s.httpHandlers.BookIsReadHandler)
	router.Path("/books/{title}").Methods(http.MethodGet).HandlerFunc(s.httpHandlers.OneBookInfoHandler)

	router.Path("/books").Methods(http.MethodGet).HandlerFunc(s.httpHandlers.ListBooksHandler)

	router.Path("/books/{title}").Methods(http.MethodDelete).HandlerFunc(s.httpHandlers.DeleteBookHandler)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
