// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package barrieroperator

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(BarrierOperatorSet{})
}

// <no value>
type BarrierOperatorSet struct {
	RequestedBarrierState byte
}

func (cmd *BarrierOperatorSet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RequestedBarrierState = payload[i]
	i++

	return nil
}

func (cmd *BarrierOperatorSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.RequestedBarrierState)

	return
}
