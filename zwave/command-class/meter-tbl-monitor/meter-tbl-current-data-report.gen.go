// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblmonitor

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(MeterTblCurrentDataReport{})
}

// <no value>
type MeterTblCurrentDataReport struct {
	ReportsToFollow byte

	Properties1 struct {
		RateType byte
	}

	Dataset uint32

	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte
}

func (cmd *MeterTblCurrentDataReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.RateType = (payload[i] & 0x03)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Dataset = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Month = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Day = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SecondLocalTime = payload[i]
	i++

	return nil
}

func (cmd *MeterTblCurrentDataReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.ReportsToFollow)

	{
		var val byte

		val |= (cmd.Properties1.RateType) & byte(0x03)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Dataset)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Year)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.Month)

	payload = append(payload, cmd.Day)

	payload = append(payload, cmd.HourLocalTime)

	payload = append(payload, cmd.MinuteLocalTime)

	payload = append(payload, cmd.SecondLocalTime)

	return
}
