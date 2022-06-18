package main  

import (
  "fmt"
  "log"
  "strings"
  "net/http"
)

func handleFormSubmit(w http.ResponseWriter, r *http.Request){
  if err := r.ParseForm(); err != nil {
    fmt.Fprintf(w, "error %v", err)
    return
  }
  
  fmt.Fprintf(w, "Post successful")
  name := r.FormValue("name")
  age := r.FormValue("age")
  
  if len(name) > 0 && (age > 0 && age <= 100) {
    fmt.Fprintf(w, "Your name submitted as %s \n", name)
    fmt.Fprintf(w, "Your age submitted as %d", age)
  } else {
    fmt.Fprintf(w, "No name/data Provided")
  }
}

func main() {  
   Fileserver := http.FileServer(http.Dir("./static"))
   http.Handle("/", Fileserver)
   http.HandleFunc("/form", handleFormSubmit)
   
   fmt.Printf("Starting server on port %d", 8080)
   if err := http.ListenAndServe(":8080", nil); err != nil {
     log.Fatal(err)
   }
}