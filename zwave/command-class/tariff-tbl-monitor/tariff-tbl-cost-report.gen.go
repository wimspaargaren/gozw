// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package tarifftblmonitor

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(TariffTblCostReport{})
}

// <no value>
type TariffTblCostReport struct {
	RateParameterSetId byte

	Properties1 struct {
		RateType byte
	}

	StartYear uint16

	StartMonth byte

	StartDay byte

	StartHourLocalTime byte

	StartMinuteLocalTime byte

	StopYear uint16

	StopMonth byte

	StopDay byte

	StopHourLocalTime byte

	StopMinuteLocalTime byte

	Currency uint32

	Properties2 struct {
		CostPrecision byte
	}

	CostValue uint32
}

func (cmd *TariffTblCostReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.RateType = (payload[i] & 0x03)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartYear = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartHourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopYear = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopHourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Currency = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.CostPrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CostValue = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	return nil
}

func (cmd *TariffTblCostReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.RateParameterSetId)

	{
		var val byte

		val |= (cmd.Properties1.RateType) & byte(0x03)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.StartYear)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.StartMonth)

	payload = append(payload, cmd.StartDay)

	payload = append(payload, cmd.StartHourLocalTime)

	payload = append(payload, cmd.StartMinuteLocalTime)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.StopYear)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.StopMonth)

	payload = append(payload, cmd.StopDay)

	payload = append(payload, cmd.StopHourLocalTime)

	payload = append(payload, cmd.StopMinuteLocalTime)

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Currency)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	{
		var val byte

		val |= (cmd.Properties2.CostPrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.CostValue)
		payload = append(payload, buf...)
	}

	return
}
