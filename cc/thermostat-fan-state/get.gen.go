// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatfanstate

import (
	"encoding/gob"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandGet cc.CommandID = 0x02

func init() {
	gob.Register(Get{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x45),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewGet)
}

func NewGet() cc.Command {
	return &Get{}
}

// <no value>
type Get struct {
}

func (cmd Get) CommandClassID() cc.CommandClassID {
	return 0x45
}

func (cmd Get) CommandID() cc.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "THERMOSTAT_FAN_STATE_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
