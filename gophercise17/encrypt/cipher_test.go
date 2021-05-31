package encrypt

import (
	"crypto/aes"
	"fmt"
	"os"
	"testing"
)
var text = "This text needs encryption"
var key = "123456789012345678901234"
var testf = "testFile.txt"
func TestHash(t *testing.T){
	fmt.Print("\nTest Hash Key\n")
	fmt.Print("\n****************************************************************\n")
	plaintext := "This is the plaintext"
	fmt.Printf("Text is- \"%s\"\n",plaintext)
	CreateHash(plaintext)
	fmt.Printf("Created Hash key is :- %x",cipherKey)
	fmt.Print("\n****************************************************************\n\n")
}

func TestEncryption(t * testing.T){
	fmt.Print("\nTest Encryption\n")
	fmt.Print("\n****************************************************************\n")
	fmt.Printf("Text is- \"%s\"\nKey is- %s\n",text,key)
	enc,err := Encryption(key,text)
	if err!= nil {
		fmt.Print(err)
	}
	fmt.Print("\nAfter applying encryption algorithm on text..\n")
	fmt.Println("Encrypted text is:-",enc)
	fmt.Print("****************************************************************\n\n")
}

func TestDecryption(t *testing.T){
	fmt.Print("\nTest Decryption\n")
	fmt.Print("\n****************************************************************\n")
	enc,_ := Encryption(key,text)
	dec,err := Decryption(key,enc)
	if err != nil {
		t.Error(err)
	}
	fmt.Print("The above encrypted text is decrypted by using decryption algorithm..\n")
	fmt.Println("Decrypted text is :-",dec)
	fmt.Print("****************************************************************\n\n")
}

func TestEncStreamWriter(t *testing.T){
	//create
	_,err := os.Stat(testf)
	if os.IsNotExist(err){
		_, err := os.Create(testf)
		if err != nil {
			fmt.Print(err)
		}
	}
	//open
	file,_ := os.Open(testf)
	defer file.Close()
	
	//write
	fmt.Print("\nTest EncryptWriter\n")
	fmt.Print("\n****************************************************************")
	_,err = EncryptW(text,file)
	if err !=nil{
		fmt.Print(err)
	}
	fmt.Print("****************************************************************\n\n")
}

func TestDecStreamReader(t *testing.T){
	
	file, err := os.Open(testf)
    if err !=nil {
        return
    }
    defer file.Close()
    text := make([]byte, 1024)
	fmt.Print("\nTest DecryptReader\n")
	fmt.Print("\n****************************************************************")
    for {
        _, err = DecryptR(string(text),file)

        if err != nil{
            fmt.Print(err)
			break
        }
    }
	fmt.Print("\n****************************************************************\n\n")
}	


func TestHexDecode(t *testing.T) {
	fmt.Print("\nTest HexDecode\n")
	fmt.Print("\n****************************************************************\n")
	_, err := Decryption("text", "1")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("\n****************************************************************\n\n")
}

func TestRand(t *testing.T){
	defer func(){
		if r :=recover(); r!=nil{
			fmt.Print(r)
		}
	}()
	fmt.Print("\nTest Invalid IV\n")
	fmt.Print("\n****************************************************************\n")
	block := aes.BlockSize+(2)
	ciphertext := make([]byte, block)
	iv := ciphertext[1:2]
	_,err := decryptStream(key,iv)
	if err!=nil{
		fmt.Print(err)
	}
}

func TestCipherSize(t *testing.T){
	fmt.Print("\n****************************************************************\n\n")
	fmt.Print("\nTest CipherSize\n")
	fmt.Print("\n****************************************************************\n")
	_, err := Decryption("text", "1123")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("\n****************************************************************\n\n")
	
}

