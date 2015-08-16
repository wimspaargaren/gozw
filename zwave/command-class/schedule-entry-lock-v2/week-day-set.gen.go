// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ScheduleEntryLockWeekDaySet{})
}

// <no value>
type ScheduleEntryLockWeekDaySet struct {
	SetAction byte

	UserIdentifier byte

	ScheduleSlotId byte

	DayOfWeek byte

	StartHour byte

	StartMinute byte

	StopHour byte

	StopMinute byte
}

func (cmd *ScheduleEntryLockWeekDaySet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SetAction = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ScheduleSlotId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DayOfWeek = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartHour = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMinute = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopHour = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMinute = payload[i]
	i++

	return nil
}

func (cmd *ScheduleEntryLockWeekDaySet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SetAction)

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.ScheduleSlotId)

	payload = append(payload, cmd.DayOfWeek)

	payload = append(payload, cmd.StartHour)

	payload = append(payload, cmd.StartMinute)

	payload = append(payload, cmd.StopHour)

	payload = append(payload, cmd.StopMinute)

	return
}
