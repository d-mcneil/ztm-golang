//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Time time.Time

type Member struct {
	name string
	id   int
}

type Book struct {
	title string
	id    int
}

type CheckOutInfo struct {
	checkOutTime Time
	checkInTime  Time
	isCheckedOut bool
	member       *Member
}

type Library struct {
	books   map[*Book]*CheckOutInfo
	members []*Member
}

func currentTime() Time {
	return Time(time.Now())
}

func checkOutBook(library *Library, book *Book, member *Member) {
	_book := library.books[book]
	_book.isCheckedOut = true
	_book.checkOutTime = currentTime()
	_book.member = member
}

func checkInBook(library *Library, book *Book) {
	_book := library.books[book]
	_book.isCheckedOut = false
	_book.checkInTime = currentTime()
}

func formatTime(t Time) string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func printLibraryInfo(library *Library) {
	fmt.Println("\n---------- Library Info ----------")

	books := make([]string, len(library.books))

	for book, info := range library.books {
		index := book.id - 1
		if info.isCheckedOut {
			books[index] = fmt.Sprintf("Book: %s was checked out by %s at %s", book.title, info.member.name, formatTime(info.checkOutTime))
			// fmt.Println("Book:", book.title, "was checked out by", info.member.name, "at", formatTime(info.checkOutTime))
		} else {
			if info.member != nil {
				books[index] = fmt.Sprintf("Book: %s is available and was checked in at %s by %s", book.title, formatTime(info.checkInTime), info.member.name)
				// fmt.Println("Book:", book.title, "is available and was checked in at", formatTime(info.checkInTime), "by", info.member.name)
			} else {
				books[index] = fmt.Sprintf("Book: %s is available", book.title)
				// fmt.Println("Book:", book.title, "is available")
			}
		}
	}

	for _, book := range books {
		fmt.Println(book)
	}
}

func main() {
	book1 := &Book{title: "Book1", id: 1}
	book2 := &Book{title: "Book2", id: 2}
	book3 := &Book{title: "Book3", id: 3}
	book4 := &Book{title: "Book4", id: 4}

	member1 := &Member{name: "Member1", id: 1}
	member2 := &Member{name: "Member2", id: 2}
	member3 := &Member{name: "Member3", id: 3}

	library := &Library{
		books: map[*Book]*CheckOutInfo{
			book1: {},
			book2: {},
			book3: {},
			book4: {},
		},
		members: []*Member{member1, member2, member3},
	}

	printLibraryInfo(library)

	checkOutBook(library, book1, member1)
	printLibraryInfo(library)

	checkOutBook(library, book2, member2)
	printLibraryInfo(library)

	checkInBook(library, book1)
	printLibraryInfo(library)

	checkOutBook(library, book1, member3)
	printLibraryInfo(library)
}
