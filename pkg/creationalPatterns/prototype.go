package creationalPatterns

import "fmt"

type Prototype[T any] interface {
	Clone() T
}

type Book struct {
	title   string
	price   float32
	content string
}

func NewBook(title string, price float32) *Book {
	return &Book{}
}

func (b *Book) Clone() *Book {
	return &Book{
		title:   b.title,
		price:   b.price,
		content: b.content + " (cached)",
	}
}

func (b *Book) fetchContentFromDb() *Book {
	b.content = "The Book Content"

	return b
}

func (b *Book) GetContent() {
	fmt.Println(b.content)
}

func ExamplePrototype() {

	original := NewBook("title", 25)
	original.fetchContentFromDb()

	clone := original.Clone()

	original.GetContent()
	clone.GetContent()
}
