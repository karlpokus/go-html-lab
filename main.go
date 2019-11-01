package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go-html-lab/internal/page"
	"go-html-lab/internal/api"
)

func main() {
	page.ParseFiles()

	router := httprouter.New()
	router.Handler("GET", "/", page.Home())
	router.Handler("GET", "/login", page.Login())

	router.Handler("POST", "/view", page.PostView())
	router.Handler("GET", "/view", api.Auth(page.GetView()))

	router.Handler("GET", "/view/:title", api.Auth(page.ViewTitle()))
	router.Handler("GET", "/edit/:title", api.Auth(page.EditTitle()))
	router.Handler("POST", "/save/:title", api.Auth(page.SaveTitle())) // TODO: move to api

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
