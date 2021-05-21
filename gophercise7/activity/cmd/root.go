package  cmd

import "github.com/spf13/cobra"

var Rcmd = &cobra.Command{
	Use:   "activity",
	Short: "Task Manager",
	Long: "It is a command line interface where you get lists of all activities.",
}
	