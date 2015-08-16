// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package alarmv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(AlarmReport{})
}

// <no value>
type AlarmReport struct {
	AlarmType byte

	AlarmLevel byte

	ZensorNetSourceNodeId byte

	ZwaveAlarmStatus byte

	ZwaveAlarmType byte

	ZwaveAlarmEvent byte

	NumberOfEventParameters byte

	EventParameter []byte
}

func (cmd *AlarmReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmLevel = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZensorNetSourceNodeId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZwaveAlarmStatus = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZwaveAlarmType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZwaveAlarmEvent = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfEventParameters = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.EventParameter = payload[i : i+6]
	i += 6

	return nil
}

func (cmd *AlarmReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.AlarmType)

	payload = append(payload, cmd.AlarmLevel)

	payload = append(payload, cmd.ZensorNetSourceNodeId)

	payload = append(payload, cmd.ZwaveAlarmStatus)

	payload = append(payload, cmd.ZwaveAlarmType)

	payload = append(payload, cmd.ZwaveAlarmEvent)

	payload = append(payload, cmd.NumberOfEventParameters)

	if cmd.EventParameter != nil && len(cmd.EventParameter) > 0 {
		payload = append(payload, cmd.EventParameter...)
	}

	return
}
