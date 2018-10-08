package arch

import (
	"fmt"
	"testing"
)

func Test_aurQuery(t *testing.T) {
	result, err := aurQuery("caddy")

	if err == EmptyResult {
		t.Log("empty")
	}

	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}
