package methods

import "time"

type Book struct {
	Title     string
	Author    string
	Pages     int
	IsRead    bool
	AddedAt   time.Time
	WasReadAt *time.Time
}

func NewBook(title string, author string, pages int) Book {
	return Book{
		Title:     title,
		Author:    author,
		Pages:     pages,
		IsRead:    false,
		AddedAt:   time.Now(),
		WasReadAt: nil,
	}

}

func (b *Book) readBook() {
	timeRead := time.Now()

	b.IsRead = true
	b.WasReadAt = &timeRead
}
func (b *Book) noReadBook() {
	timeRead := time.Now()

	b.IsRead = true
	b.WasReadAt = &timeRead
}
