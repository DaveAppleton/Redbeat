package main

import (
	"log"
	"os"
	"text/template"
)

// one OMIT

func main() {

	// Define a template.

	const letter = `
{{range .}}
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
{{end}}
`

	// two OMIT

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.

	err := t.Execute(os.Stdout, recipients)
	if err != nil {
		log.Println("executing template:", err)
	}

}

// end OMIT
