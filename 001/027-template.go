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
