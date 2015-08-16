// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv3

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ScheduleEntryLockYearDayGet{})
}

// <no value>
type ScheduleEntryLockYearDayGet struct {
	UserIdentifier byte

	ScheduleSlotId byte
}

func (cmd *ScheduleEntryLockYearDayGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

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

	return nil
}

func (cmd *ScheduleEntryLockYearDayGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.ScheduleSlotId)

	return
}
