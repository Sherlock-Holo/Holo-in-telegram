package arch

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func aurQuery(name string) (AURAnswer, error) {
	request, err := http.NewRequest(http.MethodGet, aurUrl, nil)
	if err != nil {
		return AURAnswer{}, err
	}

	query := request.URL.Query()
	query.Add("v", "5")
	query.Add("type", "search")
	query.Add("arg", name)
	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return AURAnswer{}, err
	}

	var result AurResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return AURAnswer{}, err
	}

	if result.Count == 0 {
		return AURAnswer{}, EmptyResult
	}

	res := result.Results[0]

	split := strings.Split(res.Version, "-")

	pkgrel, err := strconv.Atoi(split[len(split)-1])
	if err != nil {
		pkgrel = 1
	}

	return AURAnswer{
		OfficialAnswer: OfficialAnswer{
			Repo:    "AUR",
			Pkgname: res.Name,
			Pkgdesc: res.Desc,
			Pkgver:  res.Version,
			Pkgrel:  pkgrel,
			Url:     res.Url,
		},
		AUR: "https://aur.archlinux.org/packages/" + res.Name,
	}, nil
}
