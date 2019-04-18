package main

import (
	"os"
	"text/template"
	"fmt"
)

type Book struct{
	Title string
	Author string
}

func main() {
	//fmt.Println("Hello, playground")
	var books []Book

	
	p:= Book{Title: "Sita", Author: "Amish"}
	books = append(books, p)
	p = Book{Title: "Origin", Author: "Dan Brown"}
	books = append(books, p)
	p = Book{Title: "Fault in our stars", Author: "xxx"}
	books = append(books, p)
	
	for _, x := range books{
		fmt.Println(x.Title)
	}
	
	t := template.New("hello template")

	t, _ = t.Parse("{{range $elem := .}} The book {{$elem.Title}} was written by {{$elem.Author}} \n {{end}}")
	
	err := t.Execute(os.Stdout, books)	
	if err != nil{
		fmt.Println(err)
	}			 
	
}
