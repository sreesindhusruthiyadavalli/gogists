//File: static/main.go
package main
 
import (
  "log"
  "net/http"
)
 
func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)
 
  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
 
//open localhost:3000/example.html in your browser.
 