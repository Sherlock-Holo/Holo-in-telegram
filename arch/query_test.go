package arch

import (
	"testing"
    "fmt"
)

func TestQuery(t *testing.T) {
	answer, err := Query("caddy", "")

    if err != nil {
        t.Error(err)
    }

    fmt.Println(answer)
}
