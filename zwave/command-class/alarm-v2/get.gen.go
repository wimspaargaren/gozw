// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package alarmv2

// <no value>

type AlarmGet struct {
	AlarmType byte

	ZwaveAlarmType byte
}

func ParseAlarmGet(payload []byte) AlarmGet {
	val := AlarmGet{}

	i := 2

	val.AlarmType = payload[i]
	i++

	val.ZwaveAlarmType = payload[i]
	i++

	return val
}
