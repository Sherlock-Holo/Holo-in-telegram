package google

import (
	"errors"
	"log"
	"strings"
	"text/template"
)

const answerTplString = `<strong>Title: </strong>{{.Title}}
<strong>URL: </strong>{{.Url}}
`

var answerTpl *template.Template

var EmptyResult = errors.New("empty result")

func init() {
	answerTpl = template.Must(template.New("google answer").Parse(answerTplString))
}

type Result struct {
	Items []item `json:"items"`
}

type item struct {
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
	Link    string `json:"link"`
}

type Answer struct {
	Title string
	Url   string
}

func (a Answer) String() string {
	builder := new(strings.Builder)
	if err := answerTpl.Execute(builder, a); err != nil {
		log.Println(err)
		return ""
	}

	return builder.String()
}
