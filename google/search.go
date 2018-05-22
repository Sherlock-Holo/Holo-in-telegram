package google

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const (
    searchUrl = "https://www.googleapis.com/customsearch/v1"
)

var (
    Key string
    Cx  string
)

func Search(question string) (Answer, error) {
    request, err := http.NewRequest(http.MethodGet, searchUrl, nil)

    if err != nil {
        return Answer{}, err
    }

    query := request.URL.Query()

    query.Add("key", Key)
    query.Add("cx", Cx)
    query.Add("num", "1")
    query.Add("alt", "json")
    query.Add("q", question)

    request.URL.RawQuery = query.Encode()

    request.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

    resp, err := http.DefaultClient.Do(request)
    if err != nil {
        return Answer{}, err
    }

    /*all, err := ioutil.ReadAll(resp.Body)
      if err != nil {
          return Answer{}, err
      }*/

    // fmt.Println(string(all))

    decoder := json.NewDecoder(resp.Body)

    result := Result{}

    if err := decoder.Decode(&result); err != nil {
        return Answer{}, err
    }

    /*err = json.Unmarshal(all, &result)
      if err != nil {
          return Answer{}, err
      }*/

    if len(result.Items) == 0 {
        return Answer{}, fmt.Errorf("get 0 result")
    }

    return Answer{
        Title: result.Items[0].Title,
        Url:   result.Items[0].Link,
    }, nil
}
