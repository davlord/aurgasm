package search

import (
	"fmt"
	"sort"

	"github.com/bbrks/wrap"
	"github.com/davlord/aurgasm/common"
	"github.com/davlord/aurgasm/conf"
	u "github.com/davlord/aurgasm/util"
)

var Config *conf.Config
var wrapper wrap.Wrapper

func init() {
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

	err = markInstalled(packages)
	if err != nil {
		return err
	}
	sortPackagesByName(packages)
	printPackages(packages)

	return nil
}

func buildSearchURL(searchTerm string) string {
	return common.AurAPIBase + common.AurAPISearchPath + searchTerm
}

func aurAPISearchPackage(searchTerm string, searchResult interface{}) error {
	url := buildSearchURL(searchTerm)
	return common.GetJson(url, searchResult)
}

func sortPackagesByName(packages []*common.Package) {
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Name < packages[j].Name
	})
}

func printPackages(packages []*common.Package) {

	width, _ := u.TerminalWidth()
	colors := Config.TerminalColors()

	for _, pkg := range packages {
		fmt.Printf("%saur/%s%s %s%s%s", colors.Repo, colors.Title, pkg.Name, colors.Version, pkg.Version, colors.NoColor)
		if pkg.Installed {
			fmt.Printf(" %s[%s]%s", colors.Meta, "installed", colors.NoColor)
		}

		fmt.Printf("\n%s", wrapper.Wrap(pkg.Description, width))
	}
}

func markInstalled(packages []*common.Package) error {
	installedPackages, err := installedPackages()
	if err != nil {
		return err
	}

	for _, pkg := range packages {
		_, ok := installedPackages[pkg.Name]
		fmt.Printf("%v\n", ok)
		pkg.Installed = ok
	}

	return nil
}
