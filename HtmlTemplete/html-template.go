package main

import (
	"net/http"
	"text/template"
)

//  Here we define a person struct type that
// has Id and Name fields.

// Person ...
type Person struct {
	Id   string
	Name string
}

func renderTemplate(response http.ResponseWriter, request *http.Request) {
	person := Person{Id: "1134", Name: "Rupam Ganguly"}
	// Here we are calling ParseFiles of the html/template package, which creates a new template and
	// parses the filename we pass as an input, which is first-template.html ,
	// The resulting template will have the name and contents of the input file.

	parsTemp, _ := template.ParseFiles("./template.html")
	parsTemp.Execute(response, person)
	// 	 parsedTemplate.Execute(w, person): Here we are calling an Execute handler on a
	// parsed template, which injects person data into the template, generates an HTML
	// output, and writes it onto an HTTP response stream.
}
func main() {
	http.HandleFunc("/", renderTemplate)
	http.ListenAndServe("localhost:8080", nil)
}
