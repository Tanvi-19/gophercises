package cmd

import (
	"path/filepath"
	"testing"
	"github.com/gophercises/gophercise7/activity/database"
	"github.com/mitchellh/go-homedir"
)	
func TestAddTask(t *testing.T) {
	home,_ := homedir.Dir()
	path := filepath.Join(home,"test.db")
    database.Init(path)
	t1 := []string{"Perform some set of opepations onto db"}
	

	addCmd.Run(addCmd,t1)
	

	database.Db.Close()
}