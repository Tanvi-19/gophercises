package cmd

import (
	"fmt"

	"github.com/gophercises/gophercise7/activity/database"
	
	"github.com/spf13/cobra"
)
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "All tasks are here",
	Run: func(cmd *cobra.Command, args []string) {
		activities, _ := database.ViewAllTasks()
		
		if len(activities) == 0{
			fmt.Print("There is no task to complete\n\n")
		}
		for num, tasks := range activities{
			 fmt.Printf("%d. %s (%d) \n",num+1, tasks.Task,tasks.Num)
			 
			 
		}
	},
}

func init() {
	Rcmd.AddCommand(listCmd)
}
