// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensormultilevelv6

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SensorMultilevelGet{})
}

// <no value>
type SensorMultilevelGet struct {
	SensorType byte

	Properties1 struct {
		Scale byte
	}
}

func (cmd *SensorMultilevelGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Scale = (payload[i] & 0x18) >> 3

	i += 1

	return nil
}

func (cmd *SensorMultilevelGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Properties1.Scale << byte(3)) & byte(0x18)

		payload = append(payload, val)
	}

	return
}
