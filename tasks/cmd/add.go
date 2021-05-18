package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "To add all the tasks into the list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" task into your task \n", task)
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}