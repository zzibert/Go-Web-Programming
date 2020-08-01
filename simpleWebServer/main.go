package main

import (
  "net/http"
  "html/template"
  "math/rand"
  "time"
)

func process(w http.ResponseWriter, r *http.Request) {
  rand.Seed(time.Now().Unix())
  var t *template.Template
  if rand.Intn(10) > 5 {
    t, _ = template.ParseFiles("layout.html", "red_hello.html")
  } else {
    t, _ = template.ParseFiles("layout.html", "blue_hello.html")
  }
  t.ExecuteTemplate(w, "layout", "")
}

func form(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("form.html")
  t.Execute(w, nil)
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/process", process)
  http.HandleFunc("/form", form)
  server.ListenAndServe()
}