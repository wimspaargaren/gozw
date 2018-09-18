// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensorconfiguration

import (
	"encoding/gob"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandSensorTriggerLevelGet cc.CommandID = 0x02

func init() {
	gob.Register(SensorTriggerLevelGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x9E),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewSensorTriggerLevelGet)
}

func NewSensorTriggerLevelGet() cc.Command {
	return &SensorTriggerLevelGet{}
}

// <no value>
type SensorTriggerLevelGet struct {
}

func (cmd SensorTriggerLevelGet) CommandClassID() cc.CommandClassID {
	return 0x9E
}

func (cmd SensorTriggerLevelGet) CommandID() cc.CommandID {
	return CommandSensorTriggerLevelGet
}

func (cmd SensorTriggerLevelGet) CommandIDString() string {
	return "SENSOR_TRIGGER_LEVEL_GET"
}

func (cmd *SensorTriggerLevelGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SensorTriggerLevelGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
