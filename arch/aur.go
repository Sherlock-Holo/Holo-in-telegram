package arch

import (
	"encoding/json"
	"net/http"
)

const aurUrl = "https://aur.archlinux.org/rpc/"

type aurResult struct {
	Name    string
	Desc    string `json:"Description"`
	Version string
	Url     string `json:"URL"`
}

type AurResult struct {
	Count   int         `json:"resultcount"`
	Results []aurResult `json:"results"`
}

func aurQuery(name string) (aurResult, error) {
	request, err := http.NewRequest(http.MethodGet, aurUrl, nil)
	if err != nil {
		return aurResult{}, err
	}

	query := request.URL.Query()
	query.Add("v", "5")
	query.Add("type", "search")
	query.Add("arg", name)
	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return aurResult{}, err
	}

	var result AurResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return aurResult{}, err
	}

	if result.Count == 0 {
		return aurResult{}, EmptyResult
	}

	return result.Results[0], nil
}
