// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatmode

import (
	"encoding/gob"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandSupportedGet cc.CommandID = 0x04

func init() {
	gob.Register(SupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x40),
		Command:      cc.CommandID(0x04),
		Version:      1,
	}, NewSupportedGet)
}

func NewSupportedGet() cc.Command {
	return &SupportedGet{}
}

// <no value>
type SupportedGet struct {
}

func (cmd SupportedGet) CommandClassID() cc.CommandClassID {
	return 0x40
}

func (cmd SupportedGet) CommandID() cc.CommandID {
	return CommandSupportedGet
}

func (cmd SupportedGet) CommandIDString() string {
	return "THERMOSTAT_MODE_SUPPORTED_GET"
}

func (cmd *SupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
