package arch

import (
	"testing"
    "fmt"
)

func Test_officialQuery(t *testing.T) {
	result, err := officialQuery("linux", StableRepo...)

    if err != nil {
        t.Error(err)
    }

    fmt.Println(result)
}
