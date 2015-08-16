// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package screenmdv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ScreenMdReport{})
}

// <no value>
type ScreenMdReport struct {
	Properties1 struct {
		CharPresentation byte

		ScreenSettings byte

		MoreData bool
	}

	Properties2 struct {
		ScreenTimeout bool
	}
}

func (cmd *ScreenMdReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.CharPresentation = (payload[i] & 0x07)

	cmd.Properties1.ScreenSettings = (payload[i] & 0x38) >> 3

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.MoreData = true
	} else {
		cmd.Properties1.MoreData = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	if payload[i]&0x01 == 0x01 {
		cmd.Properties2.ScreenTimeout = true
	} else {
		cmd.Properties2.ScreenTimeout = false
	}

	i += 1

	return nil
}

func (cmd *ScreenMdReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.CharPresentation) & byte(0x07)

		val |= (cmd.Properties1.ScreenSettings << byte(3)) & byte(0x38)

		if cmd.Properties1.MoreData {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		var val byte

		if cmd.Properties2.ScreenTimeout {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}
