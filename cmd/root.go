package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/danielkvist/todots/pkg/cloner"
)

var (
	cfgFile  string
	dstRoute string
)

var rootCmd = &cobra.Command{
	Use:   "todots",
	Short: "todots is a very simple CLI that helps you to easily have a copy of all of your dotfiles",
	Long: `todots is a very simple CLI that helps you to easily have a copy of all of your dotfiles.
	
It basically copies in the directory that you specify all the files or directories that you have defined in the .todots file on your Home directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if dstRoute == "" {
			return fmt.Errorf("destination route not defined")
		}

		for k, v := range viper.AllSettings() {
			v := fmt.Sprintf("%s", v)
			finalDst := dstRoute
			if ok := strings.HasSuffix(finalDst, "/"); !ok {
				finalDst += "/"
			}

			finalDst += k + "/"
			if err := cloner.Clone(v, finalDst); err != nil {
				return err
			}
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todots.yaml)")
	rootCmd.PersistentFlags().StringVar(&dstRoute, "dst", "", "destination route")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".todots")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
