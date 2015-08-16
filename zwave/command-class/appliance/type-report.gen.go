// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package appliance

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ApplianceTypeReport{})
}

// <no value>
type ApplianceTypeReport struct {
	Properties1 struct {
		ApplianceType byte
	}

	ApplianceModeSupportedBitmask byte
}

func (cmd *ApplianceTypeReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.ApplianceType = (payload[i] & 0x3F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ApplianceModeSupportedBitmask = payload[i]
	i++

	return nil
}

func (cmd *ApplianceTypeReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.ApplianceType) & byte(0x3F)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.ApplianceModeSupportedBitmask)

	return
}
