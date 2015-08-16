// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ThermostatOperatingStateReport{})
}

// <no value>
type ThermostatOperatingStateReport struct {
	Properties1 struct {
		OperatingState byte
	}
}

func (cmd *ThermostatOperatingStateReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.OperatingState = (payload[i] & 0x0F)

	i += 1

	return nil
}

func (cmd *ThermostatOperatingStateReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.OperatingState) & byte(0x0F)

		payload = append(payload, val)
	}

	return
}
