// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SecurityNonceReport{})
}

// <no value>
type SecurityNonceReport struct {
	NonceByte []byte
}

func (cmd *SecurityNonceReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NonceByte = payload[i : i+8]

	i += 8

	return nil
}

func (cmd *SecurityNonceReport) MarshalBinary() (payload []byte, err error) {

	if paramLen := len(cmd.NonceByte); paramLen > 8 {
		return nil, errors.New("Length overflow in array parameter NonceByte")
	}

	payload = append(payload, cmd.NonceByte...)

	return
}
