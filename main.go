package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	http.HandleFunc("/todos", todosHandler)
	http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	// func ParseFiles(filenames ...string) (*Template, error) {
	// 	return parseFiles(nil, filenames...)
	// }
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, []string{"learn golang", "practise execercise", "make coffee"})

	//w.Write([]byte("hello"))
}
