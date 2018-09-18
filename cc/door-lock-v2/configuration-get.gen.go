// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package doorlockv2

import (
	"encoding/gob"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandConfigurationGet cc.CommandID = 0x05

func init() {
	gob.Register(ConfigurationGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x62),
		Command:      cc.CommandID(0x05),
		Version:      2,
	}, NewConfigurationGet)
}

func NewConfigurationGet() cc.Command {
	return &ConfigurationGet{}
}

// <no value>
type ConfigurationGet struct {
}

func (cmd ConfigurationGet) CommandClassID() cc.CommandClassID {
	return 0x62
}

func (cmd ConfigurationGet) CommandID() cc.CommandID {
	return CommandConfigurationGet
}

func (cmd ConfigurationGet) CommandIDString() string {
	return "DOOR_LOCK_CONFIGURATION_GET"
}

func (cmd *ConfigurationGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *ConfigurationGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
