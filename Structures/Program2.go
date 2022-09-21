package main

import "fmt"

type book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 book
	var Book2 book

	// Instering content of Book 1
	Book1.title = "No your story"
	Book1.author = "Savi Sharma"
	Book1.subject = "Fiction"
	Book1.book_id = 1

	// Instering content of Book 2
	Book2.title = "Payjama Profit"
	Book2.author = "Varun Maya"
	Book2.subject = "Freelancing"
	Book2.book_id = 2

	// Function to print books content
	printbook(Book1)
	printbook(Book2)
}
func printbook(b book) {
	fmt.Println("Book Title:", b.title)
	fmt.Println("Book Author:", b.author)
	fmt.Println("Book Subject:", b.subject)
	fmt.Println("Book ID:", b.book_id)
}
