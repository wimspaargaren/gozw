// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package firmwareupdatemdv3

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(FirmwareUpdateMdRequestReport{})
}

// <no value>
type FirmwareUpdateMdRequestReport struct {
	Status byte
}

func (cmd *FirmwareUpdateMdRequestReport) UnmarshalBinary(data []byte) error {
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

func (cmd *FirmwareUpdateMdRequestReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.Status)

	return
}
