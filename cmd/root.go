package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/danielkvist/todots/copier"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	dstDir  string
)

var RootCmd = &cobra.Command{
	Use:   "todots",
	Short: "todots is a simple CLI writen in Go to make a copy of your dotfiles.",
	RunE: func(cmd *cobra.Command, args []string) error {
		for k, v := range viper.AllSettings() {
			home, err := homedir.Dir()
			if err != nil {
				return fmt.Errorf("while trying to determine the HOME directory: %v", err)
			}

			srcPath := home + "/" + v.(string)
			if err := copier.Check(srcPath); err != nil {
				return fmt.Errorf("while checking source path %q: %v", srcPath, err)
			}

			dotFile := copier.NewDotfile(k)
			sf, err := os.Open(srcPath)
			if err != nil {
				return fmt.Errorf("while opening file on path %q: %v", srcPath, err)
			}
			defer sf.Close()

			r := bufio.NewReader(sf)
			if err := dotFile.CopyFrom(r); err != nil {
				return fmt.Errorf("while copying data from file %q: %v", sf.Name(), err)
			}

			_, fileName := filepath.Split(srcPath)
			dstPath := dstDir + "/" + fileName
			df, err := os.Create(dstPath)
			if err != nil {
				return fmt.Errorf("while creating destination file %q on %q: %v", fileName, dstPath, err)
			}
			defer df.Close()

			if _, err := dotFile.WriteTo(df); err != nil {
				return fmt.Errorf("while copying data from %q to destination file on %q: %v", sf.Name(), fileName, err)
			}

			if err := copier.Check(dstPath); err != nil {
				return fmt.Errorf("while checking destination path %q: %v", dstPath, err)
			}

			fmt.Printf("copy of %q successfully created.\n", fileName)
		}

		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.todots.yaml)")
	RootCmd.PersistentFlags().StringVar(&dstDir, "dst", ".", "destination directory")
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
