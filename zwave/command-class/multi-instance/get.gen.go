// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multiinstance

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(MultiInstanceGet{})
}

// <no value>
type MultiInstanceGet struct {
	CommandClass byte
}

func (cmd *MultiInstanceGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClass = payload[i]
	i++

	return nil
}

func (cmd *MultiInstanceGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.CommandClass)

	return
}
