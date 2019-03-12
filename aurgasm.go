package main

import (
	"flag"
	"fmt"

	"davlord.com/aurgasm/install"
	"davlord.com/aurgasm/search"
)

var searchTerm string
var packageName string

const (
	searchFlag  = "Ss"
	installFlag = "S"
)

func init() {
	flag.StringVar(&searchTerm, searchFlag, "", "search package")
	flag.StringVar(&packageName, installFlag, "", "install package")
}

func main() {
	flag.Parse()

	var err error

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if flagset[searchFlag] {
		err = search.SearchPackage(searchTerm)
	} else if flagset[installFlag] {
		err = install.InstallPackage(packageName)
	} else {
		flag.PrintDefaults()
	}

	if err != nil {
		fmt.Println(err)
	}
}
