// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeup

import "encoding/gob"

func init() {
	gob.Register(WakeUpNoMoreInformation{})
}

// <no value>
type WakeUpNoMoreInformation struct {
}

func (cmd *WakeUpNoMoreInformation) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *WakeUpNoMoreInformation) MarshalBinary() (payload []byte, err error) {

	return
}
