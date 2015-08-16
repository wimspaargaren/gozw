// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package time

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(DateReport{})
}

// <no value>
type DateReport struct {
	Year uint16

	Month byte

	Day byte
}

func (cmd *DateReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

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

	return nil
}

func (cmd *DateReport) MarshalBinary() (payload []byte, err error) {

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Year)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.Month)

	payload = append(payload, cmd.Day)

	return
}
