package http

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"io"
)

func GobDecoder(reader io.Reader, res interface{}) error {
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(res); err != nil {
		return wrapErr.NewWrapErr(fmt.Errorf("gob dec.Decode"), err)
	}
	return nil
}

func JsonDecoder(reader io.Reader, res interface{}) error {
	dec := json.NewDecoder(reader)
	if err := dec.Decode(res); err != nil {
		return wrapErr.NewWrapErr(fmt.Errorf("json dec.Decode"), err)
	}
	return nil
}
