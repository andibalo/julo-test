package util

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

// struct inside main struct that is referred as an interface
// needs to be registered in gob (can be registered in this init function)
// for example: YourStruct{ChildStruct interface{}} -> gob.Register(ChildStruct{})
func Serialize(obj interface{}) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	err := e.Encode(obj)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func Deserialize(str string, pointerObj interface{}) error {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}

	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)

	err = d.Decode(pointerObj)
	if err != nil {
		return err
	}

	return nil
}
