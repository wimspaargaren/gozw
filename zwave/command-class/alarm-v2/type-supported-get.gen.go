// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package alarmv2

import "encoding/gob"

func init() {
	gob.Register(AlarmTypeSupportedGet{})
}

// <no value>
type AlarmTypeSupportedGet struct {
}

func (cmd *AlarmTypeSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *AlarmTypeSupportedGet) MarshalBinary() (payload []byte, err error) {

	return
}
