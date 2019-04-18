package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"os"
)

type Book struct {
	Title  string  'json:",omitempty"' //'json:"-"'
	Author string
}

func main() {

	var books []Book

	p := Book{Title: "Sita", Author: "Amish"}
	books = append(books, p)
	p = Book{Title: "Origin", Author: "Dan Brown"}
	books = append(books, p)
	p = Book{Title: "Fault in our stars", Author: "xxx"}
	books = append(books, p)

	//Marshalling slice of structs
	b, _ := json.Marshal(books)
	fmt.Println(string(b))
		
	//Write it to a file	
	err := ioutil.WriteFile("books", b, 0644)
	if err != nil{
	fmt.Println(err)
	}

	//Read from a file
	var data []byte	
	data ,err  = ioutil.ReadFile("books")
	if err != nil{
		fmt.Println(err)
	}	

	//unmarshalling slice of structs
	var books2 []Book
	err = json.Unmarshal(data, &books2)
	if err != nil{
		fmt.Println(err)
	}	
	fmt.Printf("UnMarshhalled : %s \n", books2)
}

	// f , err := os.OpenFile("/tmp/books", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 //    if err != nil {
 //        fmt.Println(err)
 //    }
	//var data []string
	//for _, x := range books {
		//data = append(data, string(b))
	
	//}
