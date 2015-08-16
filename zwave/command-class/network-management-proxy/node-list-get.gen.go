// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementproxy

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(NodeListGet{})
}

// <no value>
type NodeListGet struct {
	SeqNo byte
}

func (cmd *NodeListGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	return nil
}

func (cmd *NodeListGet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	return
}
