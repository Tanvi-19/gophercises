package cmd

import (

	"testing"
	"github.com/gophercises/gophercise7/activity/database"
	
)

func TestList(t *testing.T){
	path := "test.db"
	database.Init(path)
	arg := []string{}
	listCmd.Run(listCmd,arg)
	database.Db.Close()
} 


