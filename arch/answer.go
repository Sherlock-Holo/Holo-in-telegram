package arch

import (
	"errors"
	"log"
	"strings"
	"text/template"
)

const (
	officialAnswerTplString = `<strong>name: </strong>{{.Pkgname}}
<strong>description: </strong>{{.Pkgdesc}}
<strong>version: </strong>{{.Pkgver}}
<strong>pkgrel: </strong>{{.Pkgrel}}
<strong>repo: </strong>{{.Repo}}
<strong>url: </strong>{{.Url}}
`

	aurAnswerTplString = `<strong>name: </strong>{{.Pkgname}}
<strong>description: </strong>{{.Pkgdesc}}
<strong>version: </strong>{{.Pkgver}}
<strong>pkgrel: </strong>{{.Pkgrel}}
<strong>repo: </strong>{{.Repo}}
<strong>url: </strong>{{.Url}}
<strong>AUR: </strong>{{.AUR}}
`
)

var (
	officialAnswerTpl *template.Template
	aurAnswerTpl      *template.Template
)

func init() {
	officialAnswerTpl = template.Must(template.New("arch answer").Parse(officialAnswerTplString))
	aurAnswerTpl = template.Must(template.New("aur answer").Parse(aurAnswerTplString))
}

type OfficialAnswer struct {
	Pkgname string
	Pkgdesc string
	Pkgver  string
	Pkgrel  int
	Repo    string
	Url     string
}

type AURAnswer struct {
	OfficialAnswer
	AUR string
}

var EmptyResult = errors.New("empty result")

func (a OfficialAnswer) String() string {
	builder := new(strings.Builder)

	if err := officialAnswerTpl.Execute(builder, a); err != nil {
		log.Println(err)
		return ""
	}

	return builder.String()
}

func (a AURAnswer) String() string {
	builder := new(strings.Builder)

	if err := aurAnswerTpl.Execute(builder, a); err != nil {
		log.Println(err)
		return ""
	}

	return builder.String()
}
