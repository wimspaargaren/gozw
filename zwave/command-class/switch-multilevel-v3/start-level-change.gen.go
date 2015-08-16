// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchmultilevelv3

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SwitchMultilevelStartLevelChange{})
}

// <no value>
type SwitchMultilevelStartLevelChange struct {
	Properties1 struct {
		IgnoreStartLevel bool

		IncDec byte

		UpDown byte
	}

	StartLevel byte

	DimmingDuration byte

	StepSize byte
}

func (cmd *SwitchMultilevelStartLevelChange) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.IncDec = (payload[i] & 0x18) >> 3

	cmd.Properties1.UpDown = (payload[i] & 0xC0) >> 6

	if payload[i]&0x20 == 0x20 {
		cmd.Properties1.IgnoreStartLevel = true
	} else {
		cmd.Properties1.IgnoreStartLevel = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartLevel = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DimmingDuration = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StepSize = payload[i]
	i++

	return nil
}

func (cmd *SwitchMultilevelStartLevelChange) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.IncDec << byte(3)) & byte(0x18)

		val |= (cmd.Properties1.UpDown << byte(6)) & byte(0xC0)

		if cmd.Properties1.IgnoreStartLevel {
			val |= byte(0x20) // flip bits on
		} else {
			val &= ^byte(0x20) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.StartLevel)

	payload = append(payload, cmd.DimmingDuration)

	payload = append(payload, cmd.StepSize)

	return
}
