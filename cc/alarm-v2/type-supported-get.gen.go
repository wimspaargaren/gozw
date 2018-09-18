// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package alarmv2

import (
	"encoding/gob"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandTypeSupportedGet cc.CommandID = 0x07

func init() {
	gob.Register(TypeSupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x71),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewTypeSupportedGet)
}

func NewTypeSupportedGet() cc.Command {
	return &TypeSupportedGet{}
}

// <no value>
type TypeSupportedGet struct {
}

func (cmd TypeSupportedGet) CommandClassID() cc.CommandClassID {
	return 0x71
}

func (cmd TypeSupportedGet) CommandID() cc.CommandID {
	return CommandTypeSupportedGet
}

func (cmd TypeSupportedGet) CommandIDString() string {
	return "ALARM_TYPE_SUPPORTED_GET"
}

func (cmd *TypeSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *TypeSupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
