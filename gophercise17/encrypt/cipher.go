package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
)


var hasher hash.Hash
var cipherKey []byte
func CreateHash(text string) (cipher.Block, error) {
	hasher = md5.New()
	fmt.Fprint(hasher, text)
	cipherKey = hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}


func Encryption(key, plaintext string) (string, error) {
	block, _ := CreateHash(key)
	
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)
	
	stream := cipher.NewCFBEncrypter(block, iv)
	
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	
	return fmt.Sprintf("%x", ciphertext), nil
}


func Decryption(key, cipherHex string) (string, error) {
	block, _ := CreateHash(key)
	ciphertext, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt error occurred:- cipher is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

func EncryptW(key string, w io.Writer) (cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)                        
	io.ReadFull(rand.Reader, iv)
	stream, _ := encryptStream(key, iv)	
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		fmt.Print("\n")
		fmt.Print("encrypt error occurred:- unable to write\n")

	}
	return cipher.StreamWriter{S: stream, W: w}, nil
}


func DecryptR(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("\nencrypt error occurred:- unable to read ")
	}
	stream, _ := decryptStream(key, iv)
	
	return &cipher.StreamReader{S: stream, R: r}, nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, _ := CreateHash(key)

	return cipher.NewCFBDecrypter(block, iv), nil
	
}

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, _ := CreateHash(key)
	
	return cipher.NewCFBEncrypter(block, iv), nil
}