package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/new", newHandler)
	http.HandleFunc("/todos/create", createHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	// func ParseFiles(filenames ...string) (*Template, error) {
	// 	return parseFiles(nil, filenames...)
	// }
	// t, err := template.ParseFiles("index.html")
	t := template.Must(template.ParseFiles("index.html", "todo.html"))

	// if err != nil {

	// log.Fatal(err)
	// }

	//t.Execute(w, []string{"learn golang", "practise execercise", "make coffee"})
	t.Execute(w, readLines("todos.txt"))
	//w.Write([]byte("hello"))
}
func newHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html", "new.html"))
	t.Execute(w, nil)

}
func createHandler(w http.ResponseWriter, r *http.Request) {
	//t := template.Must(template.ParseFiles("new.html"))
	//t.Execute(w, nil)
	//r.ParseForm()
	//log.Println(r.PostForm["item"])
	log.Println(r.FormValue("item"))

	f, err := os.OpenFile("todos.txt", os.O_APPEND, os.ModeAppend)
	check_err(err)
	_, err = fmt.Fprintln(f, r.FormValue("item"))
	check_err(err)

	http.Redirect(w, r, "/todos", http.StatusFound)

}

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}
func readLines(name string) []string {
	f, err := os.Open(name)
	if os.IsNotExist(err) {
		return nil
	}
	check_err(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check_err(scanner.Err())

	return lines
}
