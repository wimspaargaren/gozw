// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylock

import "encoding/gob"

func init() {
	gob.Register(ScheduleEntryTypeSupportedGet{})
}

// <no value>
type ScheduleEntryTypeSupportedGet struct {
}

func (cmd *ScheduleEntryTypeSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *ScheduleEntryTypeSupportedGet) MarshalBinary() (payload []byte, err error) {

	return
}
