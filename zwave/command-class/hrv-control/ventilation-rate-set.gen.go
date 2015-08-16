// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package hrvcontrol

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(HrvControlVentilationRateSet{})
}

// <no value>
type HrvControlVentilationRateSet struct {
	VentilationRate byte
}

func (cmd *HrvControlVentilationRateSet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.VentilationRate = payload[i]
	i++

	return nil
}

func (cmd *HrvControlVentilationRateSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.VentilationRate)

	return
}
