package search

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/davlord/aurgasm/common"
)

func installedPackages() (map[string]common.Package, error) {
	out, err := queryInstalledPackages()
	if err != nil {
		return nil, err
	}
	packages := parsePackages(out)
	return packages, nil
}

func parsePackages(queryResponse string) map[string]common.Package {

	packages := make(map[string]common.Package)

	scanner := bufio.NewScanner(strings.NewReader(queryResponse))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 2 {
			pkg := common.Package{
				Name:    fields[0],
				Version: fields[1],
			}
			packages[pkg.Name] = pkg
		}
	}

	return packages
}

func queryInstalledPackages() (string, error) {
	cmd := exec.Command("pacman", "-Qm")
	output, err := cmd.CombinedOutput()
	return string(output), err
}
