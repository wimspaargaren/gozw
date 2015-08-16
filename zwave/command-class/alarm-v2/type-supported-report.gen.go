// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package alarmv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(AlarmTypeSupportedReport{})
}

// <no value>
type AlarmTypeSupportedReport struct {
	Properties1 struct {
		NumberOfBitMasks byte

		V1Alarm bool
	}

	BitMask byte
}

func (cmd *AlarmTypeSupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.NumberOfBitMasks = (payload[i] & 0x1F)

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.V1Alarm = true
	} else {
		cmd.Properties1.V1Alarm = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *AlarmTypeSupportedReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.NumberOfBitMasks) & byte(0x1F)

		if cmd.Properties1.V1Alarm {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.BitMask)

	return
}
