package result

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
