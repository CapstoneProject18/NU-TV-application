package servehttp

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var IndexTemplate *template.Template

type Movies struct {
	Name string
	Path string
}

var MoviesList []Movies

func init() {
	var f *os.File
	var err error
	f, err = os.Open("mylist.txt")
	if err != nil {
		os.Exit(2121)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		MoviesList = append(MoviesList, Movies{
			Name: line,
		})
	}
	if err = scanner.Err(); err != nil {
		fmt.Print(err)
	}
	IndexTemplate, err = template.ParseFiles("IndexTemplate.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(121)
	}

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	IndexTemplate.ExecuteTemplate(w, "IndexTemplate", MoviesList)

}
func ServeHttp() {
	r := mux.NewRouter().StrictSlash(true)
	http.Handle("/static/", http.FileServer(http.Dir("static")))
	r.Handle("/", http.HandlerFunc(homeHandler))

	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
