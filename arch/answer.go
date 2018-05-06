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

var EmptyResult = fmt.Errorf("empty result")

func (a Answer) String() string {
    s := "*name*: %s" + "\n" +
        "*description*: %s" + "\n" +
        "*version*: %s" + "\n" +
        "*rel*: %d" + "\n" +
        "*repo*: %s" + "\n" +
        "*url*: %s"

    return fmt.Sprintf(s, a.Pkgname, a.Pkgdesc, a.Pkgver, a.Pkgrel, a.Repo, a.Url)
}
