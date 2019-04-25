package iostruct

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

type transform func(d []byte) ([]byte, error)

func readFile(filename string, t transform, e interface{}) error {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data, err := t(f)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(e)
}

func writeFile(filename string, t transform, data interface{}) error {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	enc.Encode(data)
	encryptedData, err := t(buf.Bytes())
	if err != nil {
		return err
	}

	f, _ := os.Create(filename)
	defer f.Close()
	_, err = f.Write(encryptedData)
	return err
}
