package main

import
(
 "encoding/json"
 "log"
 "net/http"
 
 "fmt"
 "github.com/gorilla/mux"
)

type Book struct{
	ID string
	Title string
	Author string
	Pages int
}

var books []Book


func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
    for _, item := range books {
    if item.ID == params["id"]{
    	json.NewEncoder(w).Encode(item)
    }
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = params["id"]
    books = append(books, book)
    json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	var books2 []Book
	for i, item := range books {
    if item.ID == params["id"]{
    	books = append(books[:i], books2[i+1:]...)
    	//books2 = append(books2, books2[i+1:]...)
    	break
    }
	}
	json.NewEncoder(w).Encode(books)
   
}

func main(){
	p := Book{ID: "1", Title: "Sita", Author: "Amish"}
    books = append(books, p)
    p = Book{ID:"2", Title: "Origin", Author: "Dan Brown"}
    books = append(books, p)
    p = Book{ID:"3", Title: "Fault in our stars", Author: "xxx"}
    books = append(books, p)
	router := mux.NewRouter()
	router.HandleFunc("/book", GetBooks).Methods("GET")
    router.HandleFunc("/book/{id}", GetBook).Methods("GET")
    router.HandleFunc("/book/{id}", CreateBook).Methods("POST")
    router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
