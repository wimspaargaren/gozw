// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package protectionv2

import "encoding/gob"

func init() {
	gob.Register(ProtectionSupportedGet{})
}

// <no value>
type ProtectionSupportedGet struct {
}

func (cmd *ProtectionSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *ProtectionSupportedGet) MarshalBinary() (payload []byte, err error) {

	return
}
