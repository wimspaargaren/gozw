// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package applicationcapability

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(CommandCommandClassNotSupported{})
}

// <no value>
type CommandCommandClassNotSupported struct {
	Properties1 struct {
		Dynamic bool
	}

	OffendingCommandClass byte

	OffendingCommand byte
}

func (cmd *CommandCommandClassNotSupported) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.Dynamic = true
	} else {
		cmd.Properties1.Dynamic = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.OffendingCommandClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.OffendingCommand = payload[i]
	i++

	return nil
}

func (cmd *CommandCommandClassNotSupported) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		if cmd.Properties1.Dynamic {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.OffendingCommandClass)

	payload = append(payload, cmd.OffendingCommand)

	return
}
