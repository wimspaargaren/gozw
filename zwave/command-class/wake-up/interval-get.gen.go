// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeup

import "encoding/gob"

func init() {
	gob.Register(WakeUpIntervalGet{})
}

// <no value>
type WakeUpIntervalGet struct {
}

func (cmd *WakeUpIntervalGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *WakeUpIntervalGet) MarshalBinary() (payload []byte, err error) {

	return
}
