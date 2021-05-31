package cmd

import (
	"fmt"

	"github.com/gophercises/gophercise17/encrypt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret key",
	Run: func(cmd *cobra.Command, args []string) {
		v := encrypt.File(encodingKey, vaultPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("There is no value")
			return
		}
		fmt.Printf("key is \"%s\" and value is \"%s\"\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}