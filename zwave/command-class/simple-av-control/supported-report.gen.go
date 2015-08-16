// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package simpleavcontrol

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SimpleAvControlSupportedReport{})
}

// <no value>
type SimpleAvControlSupportedReport struct {
	ReportNo byte

	BitMask byte
}

func (cmd *SimpleAvControlSupportedReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}

func (cmd *SimpleAvControlSupportedReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.ReportNo)

	payload = append(payload, cmd.BitMask)

	return
}
