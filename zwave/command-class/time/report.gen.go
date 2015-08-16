// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package time

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(TimeReport{})
}

// <no value>
type TimeReport struct {
	HourLocalTime struct {
		HourLocalTime byte

		RtcFailure bool
	}

	MinuteLocalTime byte

	SecondLocalTime byte
}

func (cmd *TimeReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HourLocalTime.HourLocalTime = (payload[i] & 0x1F)

	if payload[i]&0x80 == 0x80 {
		cmd.HourLocalTime.RtcFailure = true
	} else {
		cmd.HourLocalTime.RtcFailure = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SecondLocalTime = payload[i]
	i++

	return nil
}

func (cmd *TimeReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.HourLocalTime.HourLocalTime) & byte(0x1F)

		if cmd.HourLocalTime.RtcFailure {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.MinuteLocalTime)

	payload = append(payload, cmd.SecondLocalTime)

	return
}
