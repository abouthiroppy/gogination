package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Awesome struct {
	Items []string
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	arr := make([]string, 100)
	for i := 0; i < 100; i++ {
		arr = append(arr, "test"+strconv.Itoa(i))
	}
	awesome := Awesome{
		Items: arr,
	}
	w.Header().Set("Content-Type", "text/html")
	indexTemplate := template.Must(template.ParseFiles("view/index.html"))
	indexTemplate.Execute(w, awesome)
}
