// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package colorcontrolv2

import (
	"encoding/gob"
	"errors"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandStateGet cc.CommandID = 0x03

func init() {
	gob.Register(StateGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x33),
		Command:      cc.CommandID(0x03),
		Version:      2,
	}, NewStateGet)
}

func NewStateGet() cc.Command {
	return &StateGet{}
}

// <no value>
type StateGet struct {
	CapabilityId byte
}

func (cmd StateGet) CommandClassID() cc.CommandClassID {
	return 0x33
}

func (cmd StateGet) CommandID() cc.CommandID {
	return CommandStateGet
}

func (cmd StateGet) CommandIDString() string {
	return "STATE_GET"
}

func (cmd *StateGet) UnmarshalBinary(data []byte) error {
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

	cmd.CapabilityId = payload[i]
	i++

	return nil
}

func (cmd *StateGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.CapabilityId)

	return
}
