package packager

import (
	gonpm "github.com/nlevee/uniq-package-manager/packager/node-npm"
	gocomposer "github.com/nlevee/uniq-package-manager/packager/php-composer"

	"github.com/spf13/viper"
)

// PackageHandler itnerface
type PackageHandler interface {
	Update(path string)
	Install(path string)
}

// NewPackagerList make a new slice with packages
func NewPackagerList(config *viper.Viper) []PackageHandler {
	handlers := []PackageHandler{
		gocomposer.NewPhpComposer(config),
		gonpm.NewNodeNpm(config),
	}
	return handlers
}
