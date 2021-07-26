package http

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"io"
)

func GobEncoder(payload interface{}) (io.Reader, error) {
	if payload == nil {
		return nil, nil
	}
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(payload)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("gob enc.Encode"), err)
		return nil, err
	}
	return &network, nil
}

func JsonEncoder(payload interface{}) (io.Reader, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("json Marshal"), err)
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	return buf, nil
}
