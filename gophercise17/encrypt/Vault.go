package encrypt

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
)

func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

type Vault struct {
	encodingKey string
	filepath    string
	rout       sync.Mutex
	keyValues   map[string]string
}

func (v *Vault) loadKeyVal() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	r, _ := DecryptR(v.encodingKey, f)
	
	return v.readKeyVal(r)
}

func (v *Vault) readKeyVal(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) save() error {
	f, _ := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	w, _ := EncryptW(v.encodingKey, f)
	
	return v.writeKeyVal(w)
}

func (v *Vault) writeKeyVal(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

func (v *Vault) Get(key string) (string, error) {
	v.rout.Lock()
	defer v.rout.Unlock()
	v.loadKeyVal()
	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("there is no value for key")
	}
	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.rout.Lock()
	defer v.rout.Unlock()
	v.loadKeyVal()
	
	v.keyValues[key] = value
	err := v.save()
	return err
}