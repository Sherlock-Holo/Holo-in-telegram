package arch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	officialUrl = "https://www.archlinux.org/packages/search/json"
)

var (
	StableRepo  = []string{"Core", "Extra", "Community", "Multilib"}
	TestingRepo = []string{"Testing", "Community-Testing", "Multilib-Testing"}
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

func officialQuery(name string, repos ...string) (Answer, error) {
	request, err := http.NewRequest(http.MethodGet, officialUrl, nil)

	if err != nil {
		return Answer{}, err
	}

	query := request.URL.Query()

	query.Add("name", name)

	switch len(repos) {
	case 1:
		switch strings.ToLower(repos[0]) {
		case "", "stable":
			repos = StableRepo

		case "testing":
			repos = TestingRepo
		}

	case 0:
		repos = StableRepo
	}

	for _, repo := range repos {
		query.Add("repo", repo)
	}

	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return Answer{}, err
	}

	var result OfficialResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return Answer{}, err
	}

	if len(result.Results) == 0 {
		return Answer{}, EmptyResult
	}

	res := result.Results[0]

	pkgrel, err := strconv.Atoi(res.Rel)
	if err != nil {
		pkgrel = 1
	}

	return Answer{
		Pkgname: res.Name,
		Pkgdesc: res.Desc,
		Pkgver:  res.Version,
		Pkgrel:  pkgrel,
		Repo:    res.Repo,
		Url:     fmt.Sprintf("https://www.archlinux.org/packages/%s/%s/%s", res.Repo, res.Arch, res.Name),
	}, nil
}
