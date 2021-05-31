package cmd

import "testing"

func TestInvalidSet(t *testing.T){
	setCmd.Run(setCmd,[]string{"",""})
}
func TestSet(t *testing.T) {
	setCmd.Run(setCmd, []string{"demokey", "demoval"})
}