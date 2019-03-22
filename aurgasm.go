package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davlord/aurgasm/conf"
	"github.com/davlord/aurgasm/install"
	"github.com/davlord/aurgasm/search"
)

var searchTerm string
var packageName string
var Config *conf.Config

const (
	searchFlag  = "Ss"
	installFlag = "S"
)

func init() {
	flag.StringVar(&searchTerm, searchFlag, "", "search package")
	flag.StringVar(&packageName, installFlag, "", "install package")
}

func main() {
	initConfig()
	executeCommand()
}

func initConfig() {
	Config, err := conf.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	search.Config = Config
}

func executeCommand() {
	var err error

	flag.Parse()
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
