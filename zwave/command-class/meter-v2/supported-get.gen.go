// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package meterv2

import "encoding/gob"

func init() {
	gob.Register(MeterSupportedGet{})
}

// <no value>
type MeterSupportedGet struct {
}

func (cmd *MeterSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *MeterSupportedGet) MarshalBinary() (payload []byte, err error) {

	return
}
