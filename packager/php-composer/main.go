package gocomposer

import (
	"fmt"
	"os"
	"path"

	"github.com/docker/docker/api/types/strslice"
	"github.com/nlevee/uniq-package-manager/packager/tools"
	"github.com/spf13/viper"
)

type PhpComposer struct {
	ContainerOpts tools.ContainerOptions
}

// NewPhpComposer load viper parameters to ContainerOpts struct
func NewPhpComposer(config *viper.Viper) PhpComposer {
	c := PhpComposer{
		ContainerOpts: newDefaultOptions(),
	}

	if image := config.GetString("php-composer.image"); image != "" {
		c.ContainerOpts.Image = image
	}

	if version := config.GetString("php-composer.version"); version != "" {
		c.ContainerOpts.ImageVersion = version
	}

	return c
}

func (c PhpComposer) Update(path string) {
	if composerFile := HasComposer(path); composerFile != "" {
		if err := ComposerUpdate(composerFile, c.ContainerOpts); err != nil {
			panic(err)
		}
	}
}

func (c PhpComposer) Install(path string) {
	if composerFile := HasComposer(path); composerFile != "" {
		if err := ComposerInstall(composerFile, c.ContainerOpts); err != nil {
			panic(err)
		}
	}
}

// HasComposer check if directory given has composer.json file
func HasComposer(dir string) string {
	return tools.GetFilePath(dir, "composer.json")
}

// HasComposerLock check if directory given has composer.lock file
func HasComposerLock(dir string) string {
	return tools.GetFilePath(dir, "composer.lock")
}

// ComposerInstall must be use to install package using composer
func ComposerInstall(composerFile string, opts tools.ContainerOptions) error {
	composerPath := path.Dir(composerFile)
	fmt.Println("composer install :", composerFile, "in context", composerPath)

	if lockfile := HasComposerLock(composerPath); lockfile == "" {
		fmt.Printf("composer.lock must exist in %s\n", composerPath)
		os.Exit(1)
	}

	opts.Cmd = strslice.StrSlice{"install", "--no-scripts", "--no-progress"}
	createDockerWrapper(composerPath, opts)
	return nil
}

// ComposerUpdate must be use to update package using composer
func ComposerUpdate(composerFile string, opts tools.ContainerOptions) error {
	composerPath := path.Dir(composerFile)
	fmt.Println("composer update :", composerFile, "in context", composerPath)

	opts.Cmd = strslice.StrSlice{"update", "--no-scripts", "--no-progress"}
	createDockerWrapper(composerPath, opts)
	return nil
}
