package cmd

import (
	"fmt"
	"strings"

	"github.com/gophercises/gophercise17/encrypt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret key",
	Run: func(cmd *cobra.Command, args []string) {
		v := encrypt.File(encodingKey, vaultPath())
		key := args[0]
		value := strings.Join(args[1:], " ")
		if key == "" || value == "" {
			fmt.Println("Key and value should not be empty ...provide valid key or value")
			return
		}
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Value set successfully!.... key is \"%s\" and value is \"%s\"\n", key,value)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
