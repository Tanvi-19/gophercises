package cmd

import (
	"fmt"
	"strings"

	"github.com/gophercises/gophercise7/activity/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add tasks into the list",
	Run: func(cmd *cobra.Command, args []string) {

		
		task := strings.Join(args, " ")

		_,err := database.NewTask(task)
		if err !=nil {
			fmt.Print(err.Error()) 
		}
		fmt.Print("\nYou are trying to add following tasks into the list-\n\n")
		fmt.Print("**********************************************")
		fmt.Printf("\n\n")
			for i,task := range args{
				fmt.Printf("%d. %s\n",i+1, task)
			}
			fmt.Print("\n")
			fmt.Print("**********************************************")
			fmt.Printf("\nGood Job!,All above tasks are added into the list...")
			fmt.Print("\n\n\n")
			
	},
}

func init(){
	Rcmd.AddCommand(addCmd)
	
}