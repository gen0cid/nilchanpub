package methods

import "errors"

var ErrBookAlreadyExists = errors.New("Book Already Exists")
var ErrNoBooksByThisAuthor = errors.New("No books by this author")
var ErrNoReadedBooks = errors.New("No readed books")
var ErrReadedBooks = errors.New("All books are readed")
var ErrNoBooksByThisTitle = errors.New("No books by this title")
var ErrBookNotFound = errors.New("Book not found")
