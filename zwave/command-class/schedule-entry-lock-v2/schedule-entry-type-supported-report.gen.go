// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ScheduleEntryTypeSupportedReport{})
}

// <no value>
type ScheduleEntryTypeSupportedReport struct {
	NumberOfSlotsWeekDay byte

	NumberOfSlotsYearDay byte
}

func (cmd *ScheduleEntryTypeSupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfSlotsWeekDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfSlotsYearDay = payload[i]
	i++

	return nil
}

func (cmd *ScheduleEntryTypeSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.NumberOfSlotsWeekDay)

	payload = append(payload, cmd.NumberOfSlotsYearDay)

	return
}
