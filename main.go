package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/new", newHandler)
	http.HandleFunc("/todos/create", createHandler)
	http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	// func ParseFiles(filenames ...string) (*Template, error) {
	// 	return parseFiles(nil, filenames...)
	// }
	// t, err := template.ParseFiles("index.html")
	t := template.Must(template.ParseFiles("index.html"))

	// if err != nil {

	// log.Fatal(err)
	// }

	//t.Execute(w, []string{"learn golang", "practise execercise", "make coffee"})
	t.Execute(w, readLines)
	//w.Write([]byte("hello"))
}
func newHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("new.html"))
	t.Execute(w, nil)

}
func createHandler(w http.ResponseWriter, r *http.Request) {
	//t := template.Must(template.ParseFiles("new.html"))
	//t.Execute(w, nil)
	//r.ParseForm()
	log.Println(r.PostForm["item"])

}

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}
func readLines(fn string) []string {
	return []string{"a", "b", "c"}
}
