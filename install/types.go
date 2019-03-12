package install

import (
	"davlord.com/aurgasm/common"
)

type PackageInstall struct {
	Package  common.Package
	BuildDir string
}

type BuildStep func(pkgInstall *PackageInstall) error
