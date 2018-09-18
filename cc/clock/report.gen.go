// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package clock

import (
	"encoding/gob"
	"errors"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandReport cc.CommandID = 0x06

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x81),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	Level struct {
		Hour byte

		Weekday byte
	}

	Minute byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x81
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "CLOCK_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.Level.Hour = (payload[i] & 0x1F)

	cmd.Level.Weekday = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Minute = payload[i]
	i++

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Level.Hour) & byte(0x1F)

		val |= (cmd.Level.Weekday << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.Minute)

	return
}
