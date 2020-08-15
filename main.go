package main

import "html/template"
import "net/http"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	// ubah nama folder root
	http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":1998", nil)
}

func home(w http.ResponseWriter, r * http.Request){
	tpl.ExecuteTemplate(w, "default.gohtml", nil)
}

func about(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}