/*
Package cmd

Copyright © 2020 Nicolas Levée

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/nlevee/uniq-package-manager/packager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:     "install [BASE_PATH]",
	Aliases: []string{"ci", "i"},
	Short:   "Install all vendor dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			os.Exit(1)
		}

		appBasePath := viper.GetString("appBasePath")
		basepath := path.Join(appBasePath, args[0])
		fmt.Println("Check package to update in directory : ", basepath)

		if !path.IsAbs(basepath) {
			basepath, _ = filepath.Abs(basepath)
		}

		jobGrp := sync.WaitGroup{}
		for _, h := range packagerHandlers {
			jobGrp.Add(1)
			go func(h packager.PackageHandler) {
				h.Install(basepath)
				jobGrp.Done()
			}(h)
		}
		jobGrp.Wait()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
