// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatsetpointv3

import "encoding/gob"

func init() {
	gob.Register(ThermostatSetpointSupportedGet{})
}

// <no value>
type ThermostatSetpointSupportedGet struct {
}

func (cmd *ThermostatSetpointSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *ThermostatSetpointSupportedGet) MarshalBinary() (payload []byte, err error) {

	return
}
