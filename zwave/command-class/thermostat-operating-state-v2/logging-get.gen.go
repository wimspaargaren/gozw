// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ThermostatOperatingStateLoggingGet{})
}

// <no value>
type ThermostatOperatingStateLoggingGet struct {
	BitMask byte
}

func (cmd *ThermostatOperatingStateLoggingGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *ThermostatOperatingStateLoggingGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.BitMask)

	return
}
