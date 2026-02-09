package http

import (
	"bookshelf/methods"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type httpHandlers struct {
	bookshelf *methods.Bookshelf
}

func NewhttpHandlers(bookshelf *methods.Bookshelf) *httpHandlers {
	return &httpHandlers{
		bookshelf: bookshelf,
	}
}

func (h *httpHandlers) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	var MessageDto massageDTO
	if err := json.NewDecoder(r.Body).Decode(&MessageDto); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.toString(), http.StatusBadRequest)
		return
	}

	if err := MessageDto.validateForCreate(); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.toString(), http.StatusBadRequest)
		return
	}

	myBook := methods.NewBook(MessageDto.Title, MessageDto.Author, MessageDto.Pages)

	if err := h.bookshelf.AddBook(myBook); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, methods.ErrBookAlreadyExists) {
			http.Error(w, errDTO.toString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.toString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(myBook, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
	}
}

func (h *httpHandlers) BookIsReadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("2")

	var readDTO ReadDTO
	if err := json.NewDecoder(r.Body).Decode(&readDTO); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.toString(), http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]

	var (
		changedBook methods.Book
		err         error
	)

	if readDTO.Read {
		changedBook, err = h.bookshelf.BookIsRead(title)
	} else {
		changedBook, err = h.bookshelf.BookIsNotRead(title)
	}
	if err != nil {
		errorDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, methods.ErrBookNotFound) {
			http.Error(w, errorDTO.toString(), http.StatusNotFound)
		} else {
			http.Error(w, errorDTO.toString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(changedBook, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *httpHandlers) OneBookInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("3")

	title := mux.Vars(r)["title"]

	book, err := h.bookshelf.OneBookInfo(title)
	if err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, methods.ErrBookNotFound) {
			http.Error(w, errDTO.toString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.toString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

func (h *httpHandlers) ALLBooksInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("4")

	books := h.bookshelf.ALLBooksInfo()
	b, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *httpHandlers) ByAuthorBooksInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("5")

	Author := r.URL.Query().Get("Author")
	byAuthorBooks, err := h.bookshelf.ByAuthorBooksInfo(Author)
	if err != nil {
		errDto := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, methods.ErrNoBooksByThisAuthor) {
			http.Error(w, errDto.toString(), http.StatusBadRequest)
		} else {
			http.Error(w, errDto.toString(), http.StatusInternalServerError)
		}
		return
	}

	if Author == "" {
		errDTO := errorDTO{
			Message: "Параметр Author не задан",
			Time:    time.Now(),
		}

		b, err := json.MarshalIndent(errDTO, "", "    ")
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write(b); err != nil {
			fmt.Println("failed to write http response", err)
			return
		}
		return
	}

	b, err := json.MarshalIndent(byAuthorBooks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *httpHandlers) ByReadedBooksInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("6")

	readedBooks, err := h.bookshelf.ByReadedBooksInfo()
	if err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, methods.ErrNoReadedBooks) {
			http.Error(w, errDTO.toString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.toString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(readedBooks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

func (h *httpHandlers) ByNoReadedBooksInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("7")

	noReadedBooks, err := h.bookshelf.ByNoReadedBooksInfo()
	if err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, methods.ErrReadedBooks) {
			http.Error(w, errDTO.toString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.toString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(noReadedBooks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *httpHandlers) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("8")

	title := mux.Vars(r)["title"]
	if err := h.bookshelf.DeleteBook(title); err != nil {
		errDTO := errorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, methods.ErrNoBooksByThisTitle) {
			http.Error(w, errDTO.toString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.toString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *httpHandlers) ListBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(0)
	fmt.Println()

	qweries := r.URL.Query()
	author := qweries.Get("Author")
	isReadStr := qweries.Get("IsRead")

	if author != "" {
		h.ByAuthorBooksInfoHandler(w, r)
		return
	}
	if isReadStr == "true" {
		h.ByReadedBooksInfoHandler(w, r)
		return
	}
	if isReadStr == "false" {
		h.ByNoReadedBooksInfoHandler(w, r)
		return
	}
	h.ALLBooksInfoHandler(w, r)
}
