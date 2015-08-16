// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package configuration

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ConfigurationGet{})
}

// <no value>
type ConfigurationGet struct {
	ParameterNumber byte
}

func (cmd *ConfigurationGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ParameterNumber = payload[i]
	i++

	return nil
}

func (cmd *ConfigurationGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.ParameterNumber)

	return
}
