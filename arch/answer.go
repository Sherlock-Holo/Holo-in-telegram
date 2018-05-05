package arch

import "fmt"

type Answer struct {
    Pkgname string
    Pkgdesc string
    Pkgver  string
    Pkgrel  int
    Repo    string
    Url     string
}

var EmptyResult = fmt.Errorf("empty results")
