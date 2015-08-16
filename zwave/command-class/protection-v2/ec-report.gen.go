// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package protectionv2

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ProtectionEcReport{})
}

// <no value>
type ProtectionEcReport struct {
	NodeId byte
}

func (cmd *ProtectionEcReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	return nil
}

func (cmd *ProtectionEcReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.NodeId)

	return
}
