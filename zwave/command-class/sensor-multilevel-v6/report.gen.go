// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensormultilevelv6

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SensorMultilevelReport{})
}

// <no value>
type SensorMultilevelReport struct {
	SensorType byte

	Level struct {
		Size byte

		Scale byte

		Precision byte
	}

	SensorValue []byte
}

func (cmd *SensorMultilevelReport) UnmarshalBinary(data []byte) error {
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

	cmd.Level.Size = (payload[i] & 0x07)

	cmd.Level.Scale = (payload[i] & 0x18) >> 3

	cmd.Level.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorValue = payload[i : i+1]
	i += 1

	return nil
}

func (cmd *SensorMultilevelReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Level.Size) & byte(0x07)

		val |= (cmd.Level.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Level.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.SensorValue != nil && len(cmd.SensorValue) > 0 {
		payload = append(payload, cmd.SensorValue...)
	}

	return
}
