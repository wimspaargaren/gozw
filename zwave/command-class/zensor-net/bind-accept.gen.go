// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package zensornet

import "encoding/gob"

func init() {
	gob.Register(BindAccept{})
}

// <no value>
type BindAccept struct {
}

func (cmd *BindAccept) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *BindAccept) MarshalBinary() (payload []byte, err error) {

	return
}
