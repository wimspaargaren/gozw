// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configurationv3

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandPropertiesGet cc.CommandID = 0x0E

func init() {
	gob.Register(PropertiesGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x0E),
		Version:      3,
	}, NewPropertiesGet)
}

func NewPropertiesGet() cc.Command {
	return &PropertiesGet{}
}

// <no value>
type PropertiesGet struct {
	ParameterNumber uint16
}

func (cmd PropertiesGet) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd PropertiesGet) CommandID() cc.CommandID {
	return CommandPropertiesGet
}

func (cmd PropertiesGet) CommandIDString() string {
	return "CONFIGURATION_PROPERTIES_GET"
}

func (cmd *PropertiesGet) UnmarshalBinary(data []byte) error {
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

	return nil
}

func (cmd *PropertiesGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterNumber)
		payload = append(payload, buf...)
	}

	return
}
