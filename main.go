package main  

import (
  "fmt"
  "log"
  "net/http"
)

func main() {  
   Fileserver := http.FileServer(http.Dir("./static"))
   http.Handle("/", Fileserver)
   
   fmt.Printf("Starting server on port %d", 8080)
   if err := http.ListenAndServe(":8080", nil); err != nil {
     log.Fatal(err)
   }
}