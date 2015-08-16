// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrol

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(CapabilityReport{})
}

// <no value>
type CapabilityReport struct {
	CapabilityMask uint16
}

func (cmd *CapabilityReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CapabilityMask = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *CapabilityReport) MarshalBinary() (payload []byte, err error) {

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.CapabilityMask)
		payload = append(payload, buf...)
	}

	return
}
