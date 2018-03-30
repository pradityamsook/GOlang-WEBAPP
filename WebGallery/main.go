package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}

var (
	tmplView      = template.Must(template.New("test").ParseFiles("base.html", "test.html", "index.html"))
	tmplEdit      = template.Must(template.New("edit").ParseFiles("base.html", "edit.html", "index.html"))
	db, _         = sql.Open("sqlite3", "cache/web.db")
	creatDatabase = "create table if not exists pages (title text, body blob, timestamp text)"
)

func (p *Page) saveCache() error {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	f := "cache/" + p.Title + ".txt"
	db.Exec(creatDatabase)
	tx, _ := db.Begin()
	_, err := stmt.Exec(p.Title, p.Body, timestamp)
	tx.Commit()
	ioutil.WriteFile(f, p.Body, 0600)
	return err
}

func load(title string) (*Page, error) {
	f := title + ".txt"
	body, err := ioutil.ReadFile(f)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	tmplView.ExecuteTemplate(w, "base", p)
	//t, _ := template.ParseFiles("test.html")
	//t.Execute(w, p)
}

func save(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/test/"+title, http.StatusFound)
}

func edit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, _ := load(title)
	tmplEdit.ExecuteTemplate(w, "base", p)
	//t, _ := template.ParseFiles("edit.html")
	//t.Execute(w, p)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	p := &Page{Title: "Test", Body: []byte("Glad you come!")}
	p.save()
	http.HandleFunc("/test/", view)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/save/", save)
	http.ListenAndServe(":8000", nil)
}
