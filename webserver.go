    //Create a file named `http1.go` and copy and paste 
    //the following code into the file.
     
    package main
     
    import (
    	"io"
    	"net/http"
        "fmt"
    )
     
    func hello(w http.ResponseWriter, r *http.Request) {
    	io.WriteString(w, "Hello world!")
    }

    func name(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        io.WriteString(w, "Hello name!")
    }
     
    func main() {
    	http.HandleFunc("/", hello)
        http.HandleFunc("/name", name)
    	http.ListenAndServe(":8000", nil)
    }
    //In the terminal, execute the command `go run http1.go`,
    // then open the browser and visit `http://localhost:8000`
    			