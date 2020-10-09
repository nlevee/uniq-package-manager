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

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uniq-package-manager",
	Short: "CLI providing unique interface for most package manager",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.udm.yaml)")
	rootCmd.PersistentFlags().StringP("app-base-path", "a", "", "Application base path")

	rootCmd.PersistentFlags().String("node-image", "docker.io/library/node", "Node image for node container")
	rootCmd.PersistentFlags().String("node-version", "lts", "Node version for node container")

	rootCmd.PersistentFlags().String("php-composer-version", "1.10", "container version for php-composer package manager")
	rootCmd.PersistentFlags().String("php-composer-image", "docker.io/library/composer", "container image for php-composer package manager")

	// binding config from config file
	viper.BindPFlag("appBasePath", rootCmd.PersistentFlags().Lookup("app-base-path"))
	viper.BindPFlag("node.version", rootCmd.PersistentFlags().Lookup("node-version"))
	viper.BindPFlag("node.image", rootCmd.PersistentFlags().Lookup("node-image"))
	viper.BindPFlag("php-composer.version", rootCmd.PersistentFlags().Lookup("php-composer-version"))
	viper.BindPFlag("php-composer.image", rootCmd.PersistentFlags().Lookup("php-composer-image"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".udm" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".udm")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
