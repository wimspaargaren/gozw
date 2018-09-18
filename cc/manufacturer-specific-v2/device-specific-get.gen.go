// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package manufacturerspecificv2

import (
	"encoding/gob"
	"errors"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandDeviceSpecificGet cc.CommandID = 0x06

func init() {
	gob.Register(DeviceSpecificGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x72),
		Command:      cc.CommandID(0x06),
		Version:      2,
	}, NewDeviceSpecificGet)
}

func NewDeviceSpecificGet() cc.Command {
	return &DeviceSpecificGet{}
}

// <no value>
type DeviceSpecificGet struct {
	Properties1 struct {
		DeviceIdType byte
	}
}

func (cmd DeviceSpecificGet) CommandClassID() cc.CommandClassID {
	return 0x72
}

func (cmd DeviceSpecificGet) CommandID() cc.CommandID {
	return CommandDeviceSpecificGet
}

func (cmd DeviceSpecificGet) CommandIDString() string {
	return "DEVICE_SPECIFIC_GET"
}

func (cmd *DeviceSpecificGet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.DeviceIdType = (payload[i] & 0x07)

	i += 1

	return nil
}

func (cmd *DeviceSpecificGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.DeviceIdType) & byte(0x07)

		payload = append(payload, val)
	}

	return
}
