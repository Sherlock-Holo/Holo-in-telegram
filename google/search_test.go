package result

import (
	"testing"
    "fmt"
)

func TestSearch(t *testing.T) {
	key := "AIzaSyCn_IE6NM_ATjZ0j5vfXIFlyW-EpGs5gsU"
	cx := "006431901905483214390:i3yxhoqkzo0"
	question := "夏娜和苦力怕不得不说的故事"

    answer, err := Search(key, cx, question)

    if err != nil {
        t.Error(err)
    }

    fmt.Println(answer)
}
