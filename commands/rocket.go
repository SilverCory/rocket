package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astrocorp42/rocket/config"
	"github.com/astrocorp42/rocket/providers/heroku"
	"github.com/astrocorp42/rocket/providers/script"
	"github.com/spf13/cobra"
)

var rocketConfigPath string

func init() {
	RocketCmd.Flags().StringVarP(&rocketConfigPath, "config", "c", "", "Use the specified configuration file (and set it's directory as the working directory")
}

// RocketCmd is the rocket's root command. It's used to actually deploy
var RocketCmd = &cobra.Command{
	Use:   "rocket",
	Short: "Deploy software as fast and easily as possible",
	Long:  "Deploy software as fast and easily as possible. See https://github.com/z0mbie42/rocket",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// change working directory as the file's
		if rocketConfigPath != "" {
			dir := filepath.Dir(rocketConfigPath)
			err = os.Chdir(dir)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			rocketConfigPath = filepath.Base(rocketConfigPath)
		}

		conf, err := config.Get(rocketConfigPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("%#v\n", conf.Script)

		err = script.Deploy(conf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = heroku.Deploy(conf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}