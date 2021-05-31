package cmd

import (
	"testing"
)

var args = []string{"t-key"}
func TestInvalidGet(t *testing.T) {
	getCmd.Run(getCmd,args)
}

func TestGet(t *testing.T){
	getCmd.Run(setCmd,[]string{"getKey","getVal"})
}

