// Accessing Structure Members
package main

import "fmt"

type books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 books
	var Book2 books

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

	//Printing content of book 1
	fmt.Println("Book 1 Title:", Book1.title)
	fmt.Println("Book 1 Author:", Book1.author)
	fmt.Println("Book 1 Subject:", Book1.subject)
	fmt.Println("Book 1 Book Id:", Book1.book_id)

	//Printing content of book 2
	fmt.Println("Book 2 Title:", Book2.title)
	fmt.Println("Book 2 Author:", Book2.author)
	fmt.Println("Book 2 Subject:", Book2.subject)
	fmt.Println("Book 2 Book Id:", Book2.book_id)

}
