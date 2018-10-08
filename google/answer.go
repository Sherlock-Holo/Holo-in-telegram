package google

import "fmt"

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
	s := "*Title*: %s" + "\n" +
		"*Url*: %s"

	return fmt.Sprintf(s, a.Title, a.Url)
}
