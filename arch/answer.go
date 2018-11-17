package arch

import (
	"fmt"
	"log"
	"strings"
	"text/template"
)

const answerTplString = `<strong>name: </strong><b>Pkgname</b>
<strong>description: </strong><b>Pkgdesc</b>
<strong>version: </strong><b>Pkgver</b>
<strong>pkgrel: </strong><b>Pkgrel</b>
<strong>repo: </strong><b>Repo</b>
<strong>url: </strong><b>Url</b>
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

var EmptyResult = fmt.Errorf("empty result")

func (a Answer) String() string {
	builder := new(strings.Builder)

	if err := answerTpl.Execute(builder, a); err != nil {
		log.Println(err)
		return ""
	}

	return builder.String()
}
