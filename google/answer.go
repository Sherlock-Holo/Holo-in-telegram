package google

import "fmt"

type Result struct {
    //Kind  string `json:"kind"`
    Items []item `json:"items"`
}

type item struct {
    //Kind  string `json:"kind"`
    Title   string `json:"title"`
    Snippet string `json:"snippet"`
    Link    string `json:"link"`
}

type Answer struct {
    Title   string
    Snippet string
    Url     string
}

func (a Answer) String() string {
    s := "Title: %s" + "\n\n" +
        "Snippet: %s" + "\n\n" +
        "url: %s"

    return fmt.Sprintf(s, a.Title, a.Snippet, a.Url)
}
