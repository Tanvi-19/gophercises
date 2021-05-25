package main

import (
	"errors"
	"testing"
)
func TestMainFunc(t *testing.T){
	main()
	errMsg(errors.New("Test error"))
	
}
