// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package barrieroperator

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(BarrierOperatorSignalSupportedReport{})
}

// <no value>
type BarrierOperatorSignalSupportedReport struct {
	BitMask byte
}

func (cmd *BarrierOperatorSignalSupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *BarrierOperatorSignalSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.BitMask)

	return
}
