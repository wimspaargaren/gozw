// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package switchall

import (
	"encoding/gob"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandOff cc.CommandID = 0x05

func init() {
	gob.Register(Off{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x27),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewOff)
}

func NewOff() cc.Command {
	return &Off{}
}

// <no value>
type Off struct {
}

func (cmd Off) CommandClassID() cc.CommandClassID {
	return 0x27
}

func (cmd Off) CommandID() cc.CommandID {
	return CommandOff
}

func (cmd Off) CommandIDString() string {
	return "SWITCH_ALL_OFF"
}

func (cmd *Off) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Off) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
