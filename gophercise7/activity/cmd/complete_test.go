package cmd

import (
	
	"path/filepath"
	"testing"

	"github.com/gophercises/gophercise7/activity/database"
	"github.com/mitchellh/go-homedir"
)

func TestCompleteTask(t * testing.T){
    home, _ := homedir.Dir()
	path := filepath.Join(home, "test.db")
	database.Init(path)
	args := []string{"1","2","3","4","5","6"}
	errarg := []string{"cdh","stfdtw"}
		doCmd.Run(addCmd, args)
		doCmd.Run(doCmd,errarg)
		
}