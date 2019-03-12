package search

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"davlord.com/aurgasm/common"
)

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
	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(tw, "Name\tVersion\tDescription")
	fmt.Fprintln(tw, "----\t-------\t-----------")
	for _, pkg := range *packages {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", pkg.Name, pkg.Version, pkg.Description)
	}
	tw.Flush()
}
