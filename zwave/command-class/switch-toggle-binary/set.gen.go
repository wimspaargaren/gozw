// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchtogglebinary

import "encoding/gob"

func init() {
	gob.Register(SwitchToggleBinarySet{})
}

// <no value>
type SwitchToggleBinarySet struct {
}

func (cmd *SwitchToggleBinarySet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *SwitchToggleBinarySet) MarshalBinary() (payload []byte, err error) {

	return
}
