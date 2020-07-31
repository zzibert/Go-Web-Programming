package main

import (
  "net/http"
  "html/template"
)

func process(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("tmpl.html")
  content := `I asked: <i>whats up?</i>`
  t.Execute(w, content)
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}