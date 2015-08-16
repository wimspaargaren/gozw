// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatmodev2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ThermostatModeSet{})
}

// <no value>
type ThermostatModeSet struct {
	Level struct {
		Mode byte
	}
}

func (cmd *ThermostatModeSet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.Mode = (payload[i] & 0x1F)

	i += 1

	return nil
}

func (cmd *ThermostatModeSet) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Level.Mode) & byte(0x1F)

		payload = append(payload, val)
	}

	return
}
