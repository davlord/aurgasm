package common

import (
	"encoding/json"
	"net/http"
)

const (
	AurHost			 = "https://aur.archlinux.org"
	AurAPIBase       = AurHost+"/rpc/?v=5"
	AurAPISearchPath = "?v=5&type=search&arg="
	AurAPIInfoPath   = "?v=5&type=info&arg[]="
)

type SearchResult struct {
	Results []Package
}

type Package struct {
	Name        string
	PackageBase	string
	Version     string
	Description string
	URLPath 	string
}

func GetJson(url string, searchResult interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(searchResult)
	return nil
}
