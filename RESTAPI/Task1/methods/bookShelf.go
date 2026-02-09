package methods

import (
	"sync"
)

type Bookshelf struct {
	books map[string]Book
	mtx   sync.RWMutex
}

func NewBookShelf() *Bookshelf {
	return &Bookshelf{
		books: make(map[string]Book),
	}
}

func (b *Bookshelf) AddBook(book Book) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if _, ok := b.books[book.Title]; ok {
		return ErrBookAlreadyExists
	}

	b.books[book.Title] = book
	return nil
}

func (b *Bookshelf) BookIsRead(title string) (Book, error) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	bookCopy, ok := b.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}
	bookCopy.readBook()
	b.books[title] = bookCopy
	return b.books[title], nil
}

func (b *Bookshelf) BookIsNotRead(title string) (Book, error) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	bookCopy, ok := b.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}
	bookCopy.noReadBook()
	b.books[title] = bookCopy
	return b.books[title], nil
}

func (b *Bookshelf) OneBookInfo(title string) (Book, error) {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	bookCopy, ok := b.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	return bookCopy, nil
}

func (b *Bookshelf) ALLBooksInfo() map[string]Book {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	tmp := make(map[string]Book, len(b.books))
	for k, v := range b.books {
		tmp[k] = v
	}
	return tmp
}

func (b *Bookshelf) ByAuthorBooksInfo(author string) (map[string]Book, error) {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	tmp := make(map[string]Book)

	for k, v := range b.books {
		if v.Author == author {
			tmp[k] = v
		}
	}
	if len(tmp) == 0 {
		return tmp, ErrNoBooksByThisAuthor
	}
	return tmp, nil
}

func (b *Bookshelf) ByReadedBooksInfo() (map[string]Book, error) {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	tmp := make(map[string]Book)

	for k, v := range b.books {
		if v.IsRead == true {
			tmp[k] = v
		}
	}
	if len(tmp) == 0 {
		return tmp, ErrNoReadedBooks
	}
	return tmp, nil
}

func (b *Bookshelf) ByNoReadedBooksInfo() (map[string]Book, error) {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	tmp := make(map[string]Book)

	for k, v := range b.books {
		if v.IsRead == false {
			tmp[k] = v
		}
	}
	if len(tmp) == 0 {
		return tmp, ErrReadedBooks
	}
	return tmp, nil
}

func (b *Bookshelf) DeleteBook(title string) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	_, ok := b.books[title]

	if !ok {
		return ErrNoBooksByThisTitle
	}

	delete(b.books, title)
	return nil
}
