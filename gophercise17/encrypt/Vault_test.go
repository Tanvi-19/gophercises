package encrypt

import (
	"fmt"
	"testing"
)

var k = "key1"
var v = "Dummy value"
var fpath = "vault.txt"
var testVault = File("Key", fpath)

func TestSet(t *testing.T) {
	err := testVault.Set(k, v)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("The key is - \"%s\" and value is - \"%s\" is set \n",k,v)
}
func TestGet(t *testing.T){
	val, err := testVault.Get(k)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Get the value - \"%s\"\n",val)
}

