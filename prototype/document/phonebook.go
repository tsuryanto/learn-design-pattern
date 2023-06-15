package document

import (
	"fmt"
	"learndesignpattern/prototype"
)

type contact struct {
	name  string
	phone string
}

type PhoneBook struct {
	BookOwner string
	data      []contact
}

func NewPhoneBook(owner string) PhoneBook {
	return PhoneBook{
		BookOwner: owner,
		data: []contact{
			{name: "John", phone: "12345"},
			{name: "Doe", phone: "54321"},
		},
	}
}

func (i PhoneBook) Print() {
	fmt.Println("book of " + i.BookOwner)
	for _, contact := range i.data {
		fmt.Println(contact.name + " " + contact.phone)
	}
	fmt.Println("--------------")
}

func (i PhoneBook) Clone() prototype.Document {
	return &PhoneBook{
		BookOwner: i.BookOwner,
		data:      i.data,
	}
}
