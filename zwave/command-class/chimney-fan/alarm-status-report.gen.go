// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ChimneyFanAlarmStatusReport{})
}

// <no value>
type ChimneyFanAlarmStatusReport struct {
	AlarmStatus struct {
		NotUsed byte

		Service bool

		ExternalAlarm bool

		SensorError bool

		AlarmTemperatureExceeded bool

		SpeedChangeEnable bool

		StartTemperatureExceeded bool
	}
}

func (cmd *ChimneyFanAlarmStatusReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmStatus.NotUsed = (payload[i] & 0x30) >> 4

	if payload[i]&0x01 == 0x01 {
		cmd.AlarmStatus.Service = true
	} else {
		cmd.AlarmStatus.Service = false
	}

	if payload[i]&0x02 == 0x02 {
		cmd.AlarmStatus.ExternalAlarm = true
	} else {
		cmd.AlarmStatus.ExternalAlarm = false
	}

	if payload[i]&0x04 == 0x04 {
		cmd.AlarmStatus.SensorError = true
	} else {
		cmd.AlarmStatus.SensorError = false
	}

	if payload[i]&0x08 == 0x08 {
		cmd.AlarmStatus.AlarmTemperatureExceeded = true
	} else {
		cmd.AlarmStatus.AlarmTemperatureExceeded = false
	}

	if payload[i]&0x40 == 0x40 {
		cmd.AlarmStatus.SpeedChangeEnable = true
	} else {
		cmd.AlarmStatus.SpeedChangeEnable = false
	}

	if payload[i]&0x80 == 0x80 {
		cmd.AlarmStatus.StartTemperatureExceeded = true
	} else {
		cmd.AlarmStatus.StartTemperatureExceeded = false
	}

	i += 1

	return nil
}

func (cmd *ChimneyFanAlarmStatusReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.AlarmStatus.NotUsed << byte(4)) & byte(0x30)

		if cmd.AlarmStatus.Service {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		if cmd.AlarmStatus.ExternalAlarm {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.AlarmStatus.SensorError {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		if cmd.AlarmStatus.AlarmTemperatureExceeded {
			val |= byte(0x08) // flip bits on
		} else {
			val &= ^byte(0x08) // flip bits off
		}

		if cmd.AlarmStatus.SpeedChangeEnable {
			val |= byte(0x40) // flip bits on
		} else {
			val &= ^byte(0x40) // flip bits off
		}

		if cmd.AlarmStatus.StartTemperatureExceeded {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}
