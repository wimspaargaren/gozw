// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package applicationstatus

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(ApplicationRejectedRequest{})
}

// <no value>
type ApplicationRejectedRequest struct {
	Status byte
}

func (cmd *ApplicationRejectedRequest) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Status = payload[i]
	i++

	return nil
}

func (cmd *ApplicationRejectedRequest) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Status)

	return
}
