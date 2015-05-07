package main

import (
	"os"
	"text/template"
)

type Card struct {
	Name   string
	Title  string
	Phones []string
}

// START OMIT
var cardTemplate = template.Must(template.New("card").Parse(`
<html>
<body>
<h1>{{.Name}}</h1>
<h2>{{.Title}}</h2>
<ul>
{{range .Phones}}
	<li>{{.}}</li>
{{end}}
</ul>
</body>
</html>
`))

func main() {
	card := Card{
		Name:  "John Doe",
		Title: "software engineer",
		Phones: []string{
			"+31 62244939",
			"+31 03049430",
		},
	}
	cardTemplate.Execute(os.Stdout, card)
}

// END OMIT
