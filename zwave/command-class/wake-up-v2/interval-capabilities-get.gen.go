// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeupv2

import "encoding/gob"

func init() {
	gob.Register(WakeUpIntervalCapabilitiesGet{})
}

// <no value>
type WakeUpIntervalCapabilitiesGet struct {
}

func (cmd *WakeUpIntervalCapabilitiesGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *WakeUpIntervalCapabilitiesGet) MarshalBinary() (payload []byte, err error) {

	return
}
