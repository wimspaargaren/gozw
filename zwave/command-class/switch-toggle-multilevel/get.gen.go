// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchtogglemultilevel

import "encoding/gob"

func init() {
	gob.Register(SwitchToggleMultilevelGet{})
}

// <no value>
type SwitchToggleMultilevelGet struct {
}

func (cmd *SwitchToggleMultilevelGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *SwitchToggleMultilevelGet) MarshalBinary() (payload []byte, err error) {

	return
}
