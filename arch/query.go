package arch

import (
	"fmt"
	"strconv"
	"strings"
)

func Query(name, repo string) (Answer, error) {
	if strings.ToLower(repo) == "aur" {
		return aur(name)

	}

	answer, err := official(name, repo)
	switch err {
	case EmptyResult:
		return aur(name)
	default:
		return Answer{}, err

	case nil:
	}

	return answer, nil
}

func aur(name string) (Answer, error) {
	result, err := aurQuery(name)
	if err != nil {
		return Answer{}, err
	}

	list := strings.Split(result.Version, "-")
	pkgrel, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		pkgrel = 1
	}

	return Answer{
		Repo:    "AUR",
		Pkgname: result.Name,
		Pkgdesc: result.Desc,
		Pkgver:  result.Version,
		Pkgrel:  pkgrel,
		Url:     result.Url,
	}, nil
}

func official(name, repo string) (Answer, error) {
	var result officialResult
	var err error

	switch strings.ToLower(repo) {
	case "", "stable":
		result, err = officialQuery(name, StableRepo...)
		if err != nil {
			return Answer{}, err
		}

	case "testing":
		result, err = officialQuery(name, TestingRepo...)
		if err != nil {
			return Answer{}, err
		}

	default:
		result, err = officialQuery(name, repo)
		if err != nil {
			return Answer{}, err
		}
	}

	pkgrel, _ := strconv.Atoi(result.Rel)

	url := fmt.Sprintf("https://www.archlinux.org/packages/%s/%s/%s", result.Repo, result.Arch, result.Name)

	return Answer{
		Pkgname: result.Name,
		Pkgdesc: result.Desc,
		Pkgver:  result.Version,
		Pkgrel:  pkgrel,
		Repo:    result.Repo,
		Url:     url,
	}, nil
}
