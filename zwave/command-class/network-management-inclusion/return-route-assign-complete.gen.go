// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementinclusion

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ReturnRouteAssignComplete{})
}

// <no value>
type ReturnRouteAssignComplete struct {
	SeqNo byte

	Status byte
}

func (cmd *ReturnRouteAssignComplete) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Status = payload[i]
	i++

	return nil
}

func (cmd *ReturnRouteAssignComplete) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Status)

	return
}
