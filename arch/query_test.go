package arch

import (
    "testing"
    "fmt"
)

func TestQuery(t *testing.T) {
    answer, err := Query("linux", "")

    if err != nil {
        t.Error(err)
    }

    fmt.Println(answer)
}
