// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensorconfiguration

// <no value>

type SensorTriggerLevelReport struct {
	SensorType byte

	Size byte

	Scale byte

	Precision byte

	TriggerValue []byte
}

func ParseSensorTriggerLevelReport(payload []byte) SensorTriggerLevelReport {
	val := SensorTriggerLevelReport{}

	i := 2

	val.SensorType = payload[i]
	i++

	val.Size = (payload[i] & 0x07)

	val.Scale = (payload[i] & 0x18) << 3

	val.Precision = (payload[i] & 0xE0) << 5

	i += 1

	val.TriggerValue = payload[i : i+1]
	i += 1

	return val
}
