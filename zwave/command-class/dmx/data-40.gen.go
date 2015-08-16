// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package dmx

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(DmxData40{})
}

// <no value>
type DmxData40 struct {
	Source byte

	Properties1 struct {
		Page byte

		SequenceNo byte
	}

	DmxChannel []byte
}

func (cmd *DmxData40) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Source = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Page = (payload[i] & 0x0F)

	cmd.Properties1.SequenceNo = (payload[i] & 0x30) >> 4

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DmxChannel = payload[i : i+40]

	i += 40

	return nil
}

func (cmd *DmxData40) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Source)

	{
		var val byte

		val |= (cmd.Properties1.Page) & byte(0x0F)

		val |= (cmd.Properties1.SequenceNo << byte(4)) & byte(0x30)

		payload = append(payload, val)
	}

	if paramLen := len(cmd.DmxChannel); paramLen > 40 {
		return nil, errors.New("Length overflow in array parameter DmxChannel")
	}

	payload = append(payload, cmd.DmxChannel...)

	return
}
