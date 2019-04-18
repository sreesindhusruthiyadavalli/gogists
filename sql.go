    //sql code
    /*CREATE TABLE books (
      isbn    char(14) NOT NULL,
      title   varchar(255) NOT NULL,
      author  varchar(255) NOT NULL,
      price   decimal(5,2) NOT NULL
    );
     
    INSERT INTO books (isbn, title, author, price) VALUES
    ('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
    ('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
    ('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);
     
    ALTER TABLE books ADD PRIMARY KEY (isbn);
    */
    package main
     
    import (
    	"database/sql"
    	"fmt"
    	"log"
     
    	_ "github.com/go-sql-driver/mysql"
    )
     
    type Book struct {
    	ID   string
    	Title  string
    	Author string
    	Price  float32
    }

    
    func createBooks() []Book {
        var books []Book
        p := Book{ID: "1", Title: "Sita", Author: "Amish"}
        books = append(books, p)
        p = Book{ID:"2", Title: "Origin", Author: "Dan Brown"}
        books = append(books, p)
        p = Book{ID:"3", Title: "Fault in our stars", Author: "xxx"}
        books = append(books, p)
        return books
    }
     
    func main() {
    	db, err := sql.Open("mysql", "user11:user11@tcp(35.200.230.168:3306)/booksdb")
    	if err != nil {
    		log.Fatal(err)
    	}

        // to create a table in code
        _, err = db.Exec(`
            create table if not exists user13_books (
                id int,
                author varchar(200),
                title  varchar(200)
            );
        `)
        if err != nil {
            fmt.Println("Error creating table: ", err)
        }

        books := createBooks()
        //fmt.Println(len(books))
         
        for _, book := range books {
            fmt.Println(book)
            // simple way to create insert statement
            
            s := fmt.Sprintf(`insert into user13_books(id, author, title) values('%s','%s', '%s')`, book.ID, book.Author, book.Title)
            _, err = db.Exec(s)
            if err != nil {
            fmt.Println("Error inserting row: ", err)
            }
        }
        
    	rows, err := db.Query("SELECT * FROM user13_books")
    	if err != nil {
    		log.Fatal(err)
    	}
    	defer rows.Close()
     
    	bks := make([]*Book, 0)
    	for rows.Next() {
    		bk := new(Book)
    		err := rows.Scan(&bk.ID, &bk.Title, &bk.Author)
    		if err != nil {
    			log.Fatal(err)
    		}
    		bks = append(bks, bk)
    	}
    	if err = rows.Err(); err != nil {
    		log.Fatal(err)
    	}
     
    	for _, bk := range bks {
    		fmt.Printf("%s, %s, %s, £%.2f\n", bk.ID, bk.Title, bk.Author, bk.Price)
    	}
    }
     
    			