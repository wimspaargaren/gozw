// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package schedule

// <no value>

type CommandScheduleGet struct {
	ScheduleId byte
}

func ParseCommandScheduleGet(payload []byte) CommandScheduleGet {
	val := CommandScheduleGet{}

	i := 2

	val.ScheduleId = payload[i]
	i++

	return val
}