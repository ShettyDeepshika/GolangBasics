package main
 
import (
"fmt"
"net/http"
)
 
// CREATE A HTTP SERVER
 
// http.ResponseWriter assembles the servers response and writes to 
// the client
// http.Request is the clients request
func handler(w http.ResponseWriter, r *http.Request) {
 
// Writes to the client
fmt.Fprintf(w, "Hello World")
}
 
func handler2(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello Earth")
}
 
func main() {
 
// Calls for function handlers output to match the directory /
http.HandleFunc("/", handler)
 
// Calls for function handler2 output to match directory /earth
http.HandleFunc("/earth", handler2)
 
// Listen to port 8080 and handle requests
http.ListenAndServe(":8080", nil)
}