// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv2

// <no value>

type ScheduleEntryLockWeekDayGet struct {
	UserIdentifier byte

	ScheduleSlotId byte
}

func ParseScheduleEntryLockWeekDayGet(payload []byte) ScheduleEntryLockWeekDayGet {
	val := ScheduleEntryLockWeekDayGet{}

	i := 2

	val.UserIdentifier = payload[i]
	i++

	val.ScheduleSlotId = payload[i]
	i++

	return val
}
