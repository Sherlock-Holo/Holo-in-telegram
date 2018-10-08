package arch

import (
	"encoding/json"
	"net/http"
)

const (
	officialUrl = "https://www.archlinux.org/packages/search/json"
)

var (
	StableRepo  = []string{"Core", "Extra", "Community", "Multilib"}
	TestingRepo = []string{"Testing", "Community-Testing", "Multilib-Testing"}
	archs       = []string{"any", "x86_64"}
)

type officialResult struct {
	Name     string   `json:"pkgname"`
	Desc     string   `json:"pkgdesc"`
	Licenses []string `json:"licenses"`
	Version  string   `json:"pkgver"`
	Rel      string   `json:"pkgrel"`
	Arch     string   `json:"arch"`
	Repo     string   `json:"repo"`
}

type OfficialResult struct {
	Results []officialResult `json:"results"`
}

func officialQuery(name string, repos ...string) (officialResult, error) {
	request, err := http.NewRequest(http.MethodGet, officialUrl, nil)

	if err != nil {
		return officialResult{}, err
	}

	query := request.URL.Query()

	query.Add("name", name)

	for _, repo := range repos {
		query.Add("repo", repo)
	}

	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return officialResult{}, err
	}

	decoder := json.NewDecoder(response.Body)

	result := OfficialResult{}

	err = decoder.Decode(&result)

	if err != nil {
		return officialResult{}, err
	}

	if len(result.Results) == 0 {
		return officialResult{}, EmptyResult
	}

	return result.Results[0], nil
}
