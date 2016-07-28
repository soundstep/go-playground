package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {

	t := template.New("hello template") // create a new template with some name
	t, _ = t.Parse("Hello {{.Name}}!")  // parse some content and generate a template, which is an internal representation
	p := Person{Name: "Mary"}           // define an instance with required field
	t.Execute(os.Stdout, p)             // merge template ‘t’ with content of ‘p’

	fmt.Println()
	fmt.Println("----- generate error")

	p2 := Person2{Name: "Mary", nonExportedAgeField: "31"}
	t2 := template.New("nonexported template demo")
	t2, _ = t2.Parse("hello {{.Name}}! Age is {{.nonExportedAgeField}}.")
	err := t2.Execute(os.Stdout, p2)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	fmt.Println("----- template sets")

	fmt.Println("Load a set of templates with {{define}} clauses and execute:")
	s1, _ := template.ParseFiles("../assets/027-partials/t1.tmpl", "../assets/027-partials/t2.tmpl") //create a set of templates from many files.
	//Note that t1.tmpl is the file with contents "{{define "t_ab"}}a b{{template "t_cd"}}e f {{end}}"
	//Note that t2.tmpl is the file with contents "{{define "t_cd"}} c d {{end}}"

	s1.ExecuteTemplate(os.Stdout, "t_cd", nil) //just printing of c d
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "t_ab", nil) //execute t_ab which will include t_cd
	fmt.Println()
	s1.Execute(os.Stdout, nil) //since templates in this data structure are named, there is no default template and so it prints nothing

	fmt.Println("----- check template")

	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */")) //a valid template, so no panic with Must
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }")) // due to unmatched brace, there should be a panic here

}

// Person struct
type Person struct {
	Name string // exported field since it begins with a capital letter
}

// Person2 struct
type Person2 struct {
	Name                string // exported field since it begins with a capital letter
	nonExportedAgeField string // because it doesn't start with a capital letter
}
