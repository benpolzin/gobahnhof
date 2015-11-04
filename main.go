/*
Author: Ben Polzin
Github Author: https://github.com/benpolzin
Github Repo: https://github.com/benpolzin/gobahnhof

Version: 0.1

Thanks to:
https://golang.org/doc/articles/wiki/
http://astaxie.gitbooks.io/build-web-application-with-golang/
http://gohugo.io/templates/go-templates/
https://golang.org/doc/effective_go.html#web_server
*/

package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// create a struct named Page with Title and Body
type Page struct {
	Title string
	Body  []byte
}

// load the page based on a filename with .html extension
// perform error handling. if filename does not exist return error
func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// function to allow single viewHandler to serve up multiple templates.
// tmpl is template file name variable
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// serve up pages for viewing
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/go/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	renderTemplate(w, title, p)
}

// handle POST of file upload. referenced by manual upload and PRT
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("zipFileName")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

// handle POST of jabber file upload. referenced by manual upload and PRT
func uploadjabberHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("zipFileName")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

// handle POST of phone file upload. referenced by manual upload and PRT
func uploadphoneHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("prt_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

// web server
func main() {
	http.HandleFunc("/go/", viewHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/jabberprt/", uploadjabberHandler)
	http.HandleFunc("/phoneprt/", uploadphoneHandler)
	http.ListenAndServe(":8080", nil)
}
