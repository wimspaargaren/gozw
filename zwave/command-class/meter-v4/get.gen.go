// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package meterv4

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(MeterGet{})
}

// <no value>
type MeterGet struct {
	Properties1 struct {
		Scale byte

		RateType byte
	}

	Scale2 byte
}

func (cmd *MeterGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Scale = (payload[i] & 0x38) >> 3

	cmd.Properties1.RateType = (payload[i] & 0xC0) >> 6

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Scale2 = payload[i]
	i++

	return nil
}

func (cmd *MeterGet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.Scale << byte(3)) & byte(0x38)

		val |= (cmd.Properties1.RateType << byte(6)) & byte(0xC0)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.Scale2)

	return
}
