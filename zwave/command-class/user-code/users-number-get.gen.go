// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package usercode

import "encoding/gob"

func init() {
	gob.Register(UsersNumberGet{})
}

// <no value>
type UsersNumberGet struct {
}

func (cmd *UsersNumberGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *UsersNumberGet) MarshalBinary() (payload []byte, err error) {

	return
}
