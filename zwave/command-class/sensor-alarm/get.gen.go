// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensoralarm

// <no value>

type SensorAlarmGet struct {
	SensorType byte
}

func ParseSensorAlarmGet(payload []byte) SensorAlarmGet {
	val := SensorAlarmGet{}

	i := 2

	val.SensorType = payload[i]
	i++

	return val
}
