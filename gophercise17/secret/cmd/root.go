package cmd

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret ",
}

var encodingKey string

func vaultPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".vault")
}