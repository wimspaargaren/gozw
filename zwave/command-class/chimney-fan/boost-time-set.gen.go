// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ChimneyFanBoostTimeSet{})
}

// <no value>
type ChimneyFanBoostTimeSet struct {
	Time byte
}

func (cmd *ChimneyFanBoostTimeSet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Time = payload[i]
	i++

	return nil
}

func (cmd *ChimneyFanBoostTimeSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Time)

	return
}
