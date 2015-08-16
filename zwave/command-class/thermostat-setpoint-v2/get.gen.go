// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatsetpointv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ThermostatSetpointGet{})
}

// <no value>
type ThermostatSetpointGet struct {
	Level struct {
		SetpointType byte
	}
}

func (cmd *ThermostatSetpointGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.SetpointType = (payload[i] & 0x0F)

	i += 1

	return nil
}

func (cmd *ThermostatSetpointGet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Level.SetpointType) & byte(0x0F)

		payload = append(payload, val)
	}

	return
}
