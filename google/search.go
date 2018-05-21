package google

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const (
    searchUrl = "https://www.googleapis.com/customsearch/v1"
    key       = "AIzaSyCn_IE6NM_ATjZ0j5vfXIFlyW-EpGs5gsU"
    cx        = "010326660652388526533:m2s0xgark_u"
)

func Search(question string) (Answer, error) {

    request, err := http.NewRequest(http.MethodGet, searchUrl, nil)

    if err != nil {
        return Answer{}, err
    }

    query := request.URL.Query()

    query.Add("key", key)
    query.Add("cx", cx)
    query.Add("num", "1")
    query.Add("alt", "json")
    query.Add("q", question)

    request.URL.RawQuery = query.Encode()

    resp, err := http.DefaultClient.Do(request)

    decoder := json.NewDecoder(resp.Body)

    result := Result{}

    if err := decoder.Decode(&result); err != nil {
        return Answer{}, err
    }

    if len(result.Items) == 0 {
        return Answer{}, fmt.Errorf("get 0 result")
    }

    return Answer{
        Title: result.Items[0].Title,
        Url:   result.Items[0].Link,
    }, nil
}
