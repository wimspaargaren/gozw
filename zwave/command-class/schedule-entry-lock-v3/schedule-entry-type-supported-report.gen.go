// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv3

// <no value>

type ScheduleEntryTypeSupportedReport struct {
	NumberOfSlotsWeekDay byte

	NumberOfSlotsYearDay byte

	NumberOfSlotsDailyRepeating byte
}

func ParseScheduleEntryTypeSupportedReport(payload []byte) ScheduleEntryTypeSupportedReport {
	val := ScheduleEntryTypeSupportedReport{}

	i := 2

	val.NumberOfSlotsWeekDay = payload[i]
	i++

	val.NumberOfSlotsYearDay = payload[i]
	i++

	val.NumberOfSlotsDailyRepeating = payload[i]
	i++

	return val
}
