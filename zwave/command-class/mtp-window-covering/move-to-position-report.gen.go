// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package mtpwindowcovering

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(MoveToPositionReport{})
}

// <no value>
type MoveToPositionReport struct {
	Value byte
}

func (cmd *MoveToPositionReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Value = payload[i]
	i++

	return nil
}

func (cmd *MoveToPositionReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Value)

	return
}
