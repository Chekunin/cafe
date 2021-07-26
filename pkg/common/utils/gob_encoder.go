package utils

import (
	"bytes"
	"cafe/pkg/common/catcherr"
	"encoding/gob"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
)

func ToGobBytes(payload interface{}) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(payload)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("gob enc.Encode"), err)
		catcherr.Catch(err)
		return nil
	}
	return network.Bytes()
}
