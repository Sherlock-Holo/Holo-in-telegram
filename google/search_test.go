package google

import (
    "fmt"
    "testing"
)

func TestSearch(t *testing.T) {
    question := "夏娜和苦力怕不得不说的故事"

    answer, err := Search(question)

    if err != nil {
        t.Error(err)
    }

    fmt.Println(answer)
}
