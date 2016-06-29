package main

func main() {

}

type notExported struct {
	// this struct is visible only in this package as it starts with small letter
}

// Exported struct
type Exported struct { // variable starts with capital letter, so visible outside this package
	notExportedVariable int    // variable starts with small letter, so NOT visible outside package
	ExportedVariable    int    // variable starts with capital letter, so visible outside package
	s                   string // not exported
	S                   string // exported
}
