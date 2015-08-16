// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrolv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(StateSet{})
}

// <no value>
type StateSet struct {
	Properties1 struct {
		StateDataLength byte
	}

	DimmingDuration byte
}

func (cmd *StateSet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.StateDataLength = (payload[i] & 0x1F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DimmingDuration = payload[i]
	i++

	return nil
}

func (cmd *StateSet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.StateDataLength) & byte(0x1F)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.DimmingDuration)

	return
}
