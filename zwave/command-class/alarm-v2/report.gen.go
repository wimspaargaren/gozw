// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package alarmv2

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

func ParseAlarmReport(payload []byte) AlarmReport {
	val := AlarmReport{}

	i := 2

	val.AlarmType = payload[i]
	i++

	val.AlarmLevel = payload[i]
	i++

	val.ZensorNetSourceNodeId = payload[i]
	i++

	val.ZwaveAlarmStatus = payload[i]
	i++

	val.ZwaveAlarmType = payload[i]
	i++

	val.ZwaveAlarmEvent = payload[i]
	i++

	val.NumberOfEventParameters = payload[i]
	i++

	val.EventParameter = payload[i : i+6]
	i += 6

	return val
}
