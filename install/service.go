package install

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davlord/aurgasm/common"
)

func InstallPackage(packageName string) error {
	pkgInstall := new(PackageInstall)
	res := new(common.SearchResult)
	err := aurAPIInfoPackage(packageName, res)
	if err != nil {
		return err
	}
	if len(res.Results) < 1 {
		return errors.New("package not found")
	}
	pkgInstall.Package = *res.Results[0]

	buildSteps := []BuildStep{
		createBuildDirectory,
		gitClone,
		downloadSnapshot,
		extractSnapshot,
		install,
	}

	defer deleteBuildDirectory(pkgInstall)
	for i := range buildSteps {
		err = buildSteps[i](pkgInstall)
		if err != nil {
			return err
		}
	}

	return nil
}

func createBuildDirectory(pkgInstall *PackageInstall) (err error) {
	pkgInstall.BuildDir, err = ioutil.TempDir("", pkgInstall.Package.Name)
	return
}

func deleteBuildDirectory(pkg *PackageInstall) error {
	return os.RemoveAll(pkg.BuildDir)
}

func gitClone(pkgInstall *PackageInstall) error {
	repo := fmt.Sprintf("%s/%s.git", common.AurHost, pkgInstall.Package.PackageBase)
	return runCommand(pkgInstall.BuildDir, "git", "clone", repo)
}

func downloadSnapshot(pkgInstall *PackageInstall) error {
	url := common.AurHost + pkgInstall.Package.URLPath
	return runCommand(pkgInstall.BuildDir, "wget", url)
}

func extractSnapshot(pkgInstall *PackageInstall) error {
	snapshot := pkgInstall.Package.PackageBase + ".tar.gz"
	return runCommand(pkgInstall.BuildDir, "tar", "xzvf", snapshot)
}

func install(pkgInstall *PackageInstall) error {
	dir := pkgInstall.BuildDir + "/" + pkgInstall.Package.PackageBase
	return runCommand(dir, "makepkg", "-si")
}
