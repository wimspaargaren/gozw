// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package firmwareupdatemdv2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(FirmwareUpdateMdReport{})
}

// <no value>
type FirmwareUpdateMdReport struct {
	Properties1 struct {
		ReportNumber1 byte

		Last bool
	}

	ReportNumber2 byte

	Data []byte

	Checksum uint16
}

func (cmd *FirmwareUpdateMdReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.ReportNumber1 = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.Last = true
	} else {
		cmd.Properties1.Last = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportNumber2 = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Data = payload[i : len(payload)-2]
	i += len(payload) - 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Checksum = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *FirmwareUpdateMdReport) MarshalBinary() (payload []byte, err error) {

	{
		var val byte

		val |= (cmd.Properties1.ReportNumber1) & byte(0x7F)

		if cmd.Properties1.Last {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.ReportNumber2)

	payload = append(payload, cmd.Data...)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Checksum)
		payload = append(payload, buf...)
	}

	return
}
