// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package tariffconfig

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(TariffTblSet{})
}

// <no value>
type TariffTblSet struct {
	RateParameterSetId byte

	Properties1 struct {
		TariffPrecision byte
	}

	TariffValue uint32
}

func (cmd *TariffTblSet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RateParameterSetId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.TariffPrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.TariffValue = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	return nil
}

func (cmd *TariffTblSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.RateParameterSetId)

	{
		var val byte

		val |= (cmd.Properties1.TariffPrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.TariffValue)
		payload = append(payload, buf...)
	}

	return
}
