package install

import (
	"github.com/davlord/aurgasm/common"
)

type PackageInstall struct {
	Package  common.Package
	BuildDir string
}

type BuildStep func(pkgInstall *PackageInstall) error
