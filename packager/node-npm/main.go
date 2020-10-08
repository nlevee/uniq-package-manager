package gonpm

import (
	"fmt"
	"os"
	"path"

	"github.com/docker/docker/api/types/strslice"
	"github.com/nlevee/uniq-package-manager/packager/tools"
)

type NodeNpm struct{}

func (c NodeNpm) Update(path string) {
	if composerFile := HasPackage(path); composerFile != "" {
		if err := NpmUpdate(composerFile); err != nil {
			panic(err)
		}
	}
}

func (c NodeNpm) Install(path string) {
	if composerFile := HasPackage(path); composerFile != "" {
		if err := NpmInstall(composerFile); err != nil {
			panic(err)
		}
	}
}

// HasPackages check if directory given has package.json file
func HasPackage(dir string) string {
	return tools.GetFilePath(dir, "package.json")
}

// HasPackagesLock check if directory given has package-lock.json file
func HasPackageLock(dir string) string {
	return tools.GetFilePath(dir, "package-lock.json")
}

// NpmInstall must be use to install package using npm
func NpmInstall(packageFile string) error {
	packagePath := path.Dir(packageFile)
	fmt.Println("npm ci :", packageFile, "in context", packagePath)

	if lockfile := HasPackageLock(packagePath); lockfile == "" {
		fmt.Printf("package-lock.json must exist in %s\n", packagePath)
		os.Exit(1)
	}

	createDockerWrapper(packagePath, strslice.StrSlice{"ci"})
	return nil
}

// NpmUpdate must be use to update package using npm
func NpmUpdate(packageFile string) error {
	packagePath := path.Dir(packageFile)
	fmt.Println("npm install :", packageFile, "in context", packagePath)

	createDockerWrapper(packagePath, strslice.StrSlice{"install"})
	return nil
}
