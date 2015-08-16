// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensormultilevelv5

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SensorMultilevelSupportedGetScale{})
}

// <no value>
type SensorMultilevelSupportedGetScale struct {
	SensorType byte
}

func (cmd *SensorMultilevelSupportedGetScale) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorType = payload[i]
	i++

	return nil
}

func (cmd *SensorMultilevelSupportedGetScale) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SensorType)

	return
}
