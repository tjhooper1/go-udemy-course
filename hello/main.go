package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Heading string
	Input string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	
	home := Page{
		Title: "Nothing escaped",
		Heading: "Nothing escaped some more",
		Input: "whatever",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
