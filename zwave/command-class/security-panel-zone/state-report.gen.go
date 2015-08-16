// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package securitypanelzone

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SecurityPanelZoneStateReport{})
}

// <no value>
type SecurityPanelZoneStateReport struct {
	ZoneNumber byte

	ZoneState byte
}

func (cmd *SecurityPanelZoneStateReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZoneNumber = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZoneState = payload[i]
	i++

	return nil
}

func (cmd *SecurityPanelZoneStateReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.ZoneNumber)

	payload = append(payload, cmd.ZoneState)

	return
}
