// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package firmwareupdatemdv3

import "encoding/gob"

func init() {
	gob.Register(FirmwareMdGet{})
}

// <no value>
type FirmwareMdGet struct {
}

func (cmd *FirmwareMdGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *FirmwareMdGet) MarshalBinary() (payload []byte, err error) {

	return
}
