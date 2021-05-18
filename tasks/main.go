package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gophercises/tasks/cmd"
	"github.com/gophercises/tasks/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home,_ := homedir.Dir()
	path := filepath.Join(home,"tasks.db")
    must(db.Init(path))
	fmt.Println("DB CONNECT")
	must(cmd.RootCmd.Execute())
}

func must(err error){
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
