// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package meterv4

import "encoding/gob"

func init() {
	gob.Register(MeterReset{})
}

// <no value>
type MeterReset struct {
}

func (cmd *MeterReset) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *MeterReset) MarshalBinary() (payload []byte, err error) {

	return
}
