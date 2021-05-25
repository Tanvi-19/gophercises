package cmd

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/gophercises/gophercise7/activity/database"
	"github.com/mitchellh/go-homedir"
)


func TestList(t *testing.T){
	f()
	home,_ := homedir.Dir()
	path := filepath.Join(home,"demo.db")
	database.Init(path)
	
} 

func f(){
	defer func(){
		if r:= recover(); r!=nil{
			fmt.Print("recovered")
		}
	}()
	g()
}

func g(){
	arg :=[]string{}
	listCmd.Run(listCmd,arg)
	database.Db.Close()
}



