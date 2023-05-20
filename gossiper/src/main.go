package main

import (
	"net/http"
  "fmt"
  "log"
)

 func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })

    fmt.Printf("Starting server at port 8070\n")
    if err := http.ListenAndServe(":8070", nil); err != nil {
        log.Fatal(err)
    }
}
