package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/danielkvist/todots/pkg/copier"
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
			return fmt.Errorf("no destination route provided")
		}

		var err error
		if !filepath.IsAbs(dstRoute) {
			dstRoute, err = filepath.Abs(dstRoute)
			if err != nil {
				return fmt.Errorf("while making the destination route absolute: %v", err)
			}
		}

		dst := dstRoute
		if ok := strings.HasSuffix(dstRoute, "/"); !ok {
			dst += "/"
		}

		for _, k := range viper.AllKeys() {
			err := copier.Copy(fmt.Sprintf("%s", viper.Get(k)), dst+k+"/")
			if err != nil {
				return fmt.Errorf("%v", err)
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
