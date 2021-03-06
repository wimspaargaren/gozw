// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configurationv3

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandInfoReport cc.CommandID = 0x0D

func init() {
	gob.Register(InfoReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x0D),
		Version:      3,
	}, NewInfoReport)
}

func NewInfoReport() cc.Command {
	return &InfoReport{}
}

// <no value>
type InfoReport struct {
	ParameterNumber uint16

	ReportsToFollow byte

	Info []byte
}

func (cmd InfoReport) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd InfoReport) CommandID() cc.CommandID {
	return CommandInfoReport
}

func (cmd InfoReport) CommandIDString() string {
	return "CONFIGURATION_INFO_REPORT"
}

func (cmd *InfoReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ParameterNumber = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return nil
	}

	cmd.Info = payload[i:]

	return nil
}

func (cmd *InfoReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterNumber)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.ReportsToFollow)

	payload = append(payload, cmd.Info...)

	return
}
