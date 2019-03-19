package search

import (
	"fmt"
	"sort"

	"github.com/bbrks/wrap"
	"github.com/davlord/aurgasm/common"
	u "github.com/davlord/aurgasm/util"
)

var colors u.Colors
var wrapper wrap.Wrapper

func init() {
	colors = u.TerminalColors()

	wrapper = wrap.NewWrapper()
	wrapper.OutputLinePrefix = "    "
}

func SearchPackage(searchTerm string) error {
	res := new(common.SearchResult)
	err := aurAPISearchPackage(searchTerm, &res)
	if err != nil {
		return err
	}

	packages := res.Results
	sortPackagesByName(&packages)
	printPackages(&packages)

	return nil
}

func buildSearchURL(searchTerm string) string {
	return common.AurAPIBase + common.AurAPISearchPath + searchTerm
}

func aurAPISearchPackage(searchTerm string, searchResult interface{}) error {
	url := buildSearchURL(searchTerm)
	return common.GetJson(url, searchResult)
}

func sortPackagesByName(packages *[]common.Package) {
	sort.Slice(*packages, func(i, j int) bool {
		return (*packages)[i].Name < (*packages)[j].Name
	})
}

func printPackages(packages *[]common.Package) {

	width, _ := u.TerminalWidth()

	for _, pkg := range *packages {
		fmt.Printf("%saur/%s%s %s%s%s\n", colors.Repo, colors.Title, pkg.Name, colors.Version, pkg.Version, colors.NoColor)
		fmt.Printf("%s", wrapper.Wrap(pkg.Description, width))
	}
}
