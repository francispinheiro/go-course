package main

import (
	"net/http"

	"github.com/francispinheiro/go-course/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	//renderTemplate(w, "about.page.tmpl")
}
