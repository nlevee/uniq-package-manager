package packager

import (
	gonpm "github.com/nlevee/uniq-package-manager/packager/node-npm"
	gocomposer "github.com/nlevee/uniq-package-manager/packager/php-composer"
)

// PackageHandler
type PackageHandler interface {
	Update(path string)
	Install(path string)
}

func NewPackagerList() []PackageHandler {
	handlers := []PackageHandler{
		gocomposer.PhpComposer{},
		gonpm.NodeNpm{},
	}
	return handlers
}
