package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "It marks task as a complete",
	Run: func(cmd *cobra.Command, args []string) {
		var nums []int
		for _, arg := range args{
			num, err := strconv.Atoi(arg)
			if err != nil{
				fmt.Println("Error while parsing ",arg)
			}else{
				nums = append(nums, num)
			}
		}
		fmt.Println(nums)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
