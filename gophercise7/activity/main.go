package main

import (
	"fmt"
	"path/filepath"
	"github.com/gophercises/gophercise7/activity/cmd"
	"github.com/gophercises/gophercise7/activity/database"    
	"github.com/mitchellh/go-homedir"
)

func main() {
	home,_ := homedir.Dir()
	path := filepath.Join(home,"tasks.db")
    errMsg(database.Init(path))
	fmt.Println("DB CONNECT")
	errMsg(cmd.Rcmd.Execute())
}

func errMsg(err error){
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
