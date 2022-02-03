package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
