package gonpm

import (
	"fmt"
	"os"
	"path"

	"github.com/docker/docker/api/types/strslice"
	"github.com/nlevee/uniq-package-manager/packager/tools"
	"github.com/spf13/viper"
)

// NodeNpm struct
type NodeNpm struct {
	ContainerOpts tools.ContainerOptions
}

// NewNodeNpm load viper parameters to ContainerOpts struct
func NewNodeNpm(config *viper.Viper) NodeNpm {
	c := NodeNpm{
		ContainerOpts: newDefaultOptions(),
	}

	if image := config.GetString("node.image"); image != "" {
		c.ContainerOpts.Image = image
	}

	if version := config.GetString("node.version"); version != "" {
		c.ContainerOpts.ImageVersion = version
	}

	return c
}

// Update must be use to update package using npm
func (c NodeNpm) Update(path string) {
	if composerFile := HasPackage(path); composerFile != "" {
		if err := NpmUpdate(composerFile, c.ContainerOpts); err != nil {
			panic(err)
		}
	}
}

// Install must be use to install package using npm
func (c NodeNpm) Install(path string) {
	if composerFile := HasPackage(path); composerFile != "" {
		if err := NpmInstall(composerFile, c.ContainerOpts); err != nil {
			panic(err)
		}
	}
}

// HasPackage check if directory given has package.json file
func HasPackage(dir string) string {
	return tools.GetFilePath(dir, "package.json")
}

// HasPackageLock check if directory given has package-lock.json file
func HasPackageLock(dir string) string {
	return tools.GetFilePath(dir, "package-lock.json")
}

// NpmInstall must be use to install package using npm
func NpmInstall(packageFile string, opts tools.ContainerOptions) error {
	packagePath := path.Dir(packageFile)
	fmt.Println("npm ci :", packageFile, "in context", packagePath)

	if lockfile := HasPackageLock(packagePath); lockfile == "" {
		fmt.Printf("package-lock.json must exist in %s\n", packagePath)
		os.Exit(1)
	}

	opts.Cmd = strslice.StrSlice{"ci"}
	createDockerWrapper(packagePath, opts)
	return nil
}

// NpmUpdate must be use to update package using npm
func NpmUpdate(packageFile string, opts tools.ContainerOptions) error {
	packagePath := path.Dir(packageFile)
	fmt.Println("npm install :", packageFile, "in context", packagePath)

	opts.Cmd = strslice.StrSlice{"install"}
	createDockerWrapper(packagePath, opts)
	return nil
}
