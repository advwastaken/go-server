package main  

import (
  "fmt"
  "log"
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
  fmt.Fprintf(w, "Your name submitted as %s \n", name)
  fmt.Fprintf(w, "Your age submitted as %d", age)
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