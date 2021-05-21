package cmd

import (
	"fmt"
	"strconv"

	"github.com/gophercises/gophercise7/activity/database"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete is use to mark task as done",
	Run: func(cmd *cobra.Command, args []string) {
		activities, _ := database.ViewAllTasks()
		var nums []int
		for _, num := range args{
			result, err := strconv.Atoi(num)
			
			if err != nil{
				fmt.Println("\nSome error occured while parsing this:-",num)
				
			}else{
				nums = append(nums, result)
			}
		}
		for _, num := range nums{
			
			if num <= 0 || num > len(activities){	
				fmt.Println()
				fmt.Printf("Wrong task numbers:- %d",num)
				
				continue
			}
			
			activity := activities[num-1]
			
			err := database.DelTask(activity.Num)
			
			if err == nil{
				fmt.Printf("\nSuccessfully completed task-")
				fmt.Print("\n\n")
				fmt.Printf(" %d. %s ",num+1,activity.Task)
				fmt.Print("\n**********************************************")
				fmt.Print("\n\n")
				
			}else{
				fmt.Printf("%s enable to done",activity.Task)
			}
		}
		fmt.Print("\n\n")
        fmt.Println("**********************************************")
		fmt.Print("\n")
	},
	
}

func init() {
	Rcmd.AddCommand(doCmd)

}
