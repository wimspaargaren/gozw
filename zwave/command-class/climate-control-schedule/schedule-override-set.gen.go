// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package climatecontrolschedule

// <no value>

type ScheduleOverrideSet struct {
	OverrideType byte

	OverrideState byte
}

func ParseScheduleOverrideSet(payload []byte) ScheduleOverrideSet {
	val := ScheduleOverrideSet{}

	i := 2

	val.OverrideType = (payload[i] & 0x03)

	i += 1

	val.OverrideState = payload[i]
	i++

	return val
}
