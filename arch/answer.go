package arch

import (
	"errors"
	"log"
	"strings"
	"text/template"
)

const answerTplString = `<strong>name: </strong>{{.Pkgname}}
<strong>description: </strong>{{.Pkgdesc}}
<strong>version: </strong>{{.Pkgver}}
<strong>pkgrel: </strong>{{.Pkgrel}}
<strong>repo: </strong>{{.Repo}}
<strong>url: </strong>{{.Url}}
`

var answerTpl *template.Template

func init() {
	answerTpl = template.Must(template.New("arch answer").Parse(answerTplString))
}

type Answer struct {
	Pkgname string
	Pkgdesc string
	Pkgver  string
	Pkgrel  int
	Repo    string
	Url     string
}

var EmptyResult = errors.New("empty result")

func (a Answer) String() string {
	builder := new(strings.Builder)

	if err := answerTpl.Execute(builder, a); err != nil {
		log.Println(err)
		return ""
	}

	return builder.String()
}
