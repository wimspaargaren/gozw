// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import "encoding/gob"

func init() {
	gob.Register(ChimneyFanAlarmTempGet{})
}

// <no value>
type ChimneyFanAlarmTempGet struct {
}

func (cmd *ChimneyFanAlarmTempGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *ChimneyFanAlarmTempGet) MarshalBinary() (payload []byte, err error) {

	return
}
