package result

import (
    "net/http"
    "fmt"
    "encoding/json"
)

//const searchUrl = "https://www.googleapis.com/customsearch/v1?key=AIzaSyCn_IE6NM_ATjZ0j5vfXIFlyW-EpGs5gsU&cx=006431901905483214390:i3yxhoqkzo0&num=1&alt=json&q=%E6%94%AF%E4%BB%98%E5%AE%9D"
const searchUrl = "https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&num=1&alt=json&q=%s"

func Search(key, cx, question string) (Answer, error) {
    resp, err := http.Get(fmt.Sprintf(searchUrl, key, cx, question))

    if err != nil {
        return Answer{}, err
    }

    decoder := json.NewDecoder(resp.Body)

    result := Result{}

    if err := decoder.Decode(&result); err != nil {
        return Answer{}, err
    }

    if len(result.Items) == 0 {
        return Answer{}, fmt.Errorf("get 0 result")
    }

    return Answer{
        Title:   result.Items[0].Title,
        Snippet: result.Items[0].Snippet,
        Url:     result.Items[0].Link,
    }, nil
}
