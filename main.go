// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// param returns a named parameter from the request context
func param(s string, r *http.Request) string {
	return httprouter.ParamsFromContext(r.Context()).ByName(s)
}

func view() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := param("title", r)
		p, err := loadPage(title)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Redirect(w, r, "/edit/"+title, http.StatusFound)
			return
		}
		renderTemplate(w, "view", p)
	}
}

func edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := param("title", r)
		p, err := loadPage(title)
		if err != nil {
			p = &Page{Title: title} // create a page on the fly
		}
		renderTemplate(w, "edit", p)
	}
}

func save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := param("title", r)
		body := r.FormValue("body")
		p := &Page{Title: title, Body: []byte(body)}
		err := p.save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := httprouter.New()
	router.Handler("GET", "/view/:title", view())
	router.Handler("GET", "/edit/:title", edit())
	router.Handler("POST", "/save/:title", save())

	log.Fatal(http.ListenAndServe(":8080", router))
}
