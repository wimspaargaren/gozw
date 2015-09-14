// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SchemeReport{})
}

// <no value>
type SchemeReport struct {
	SupportedSecuritySchemes byte
}

func (cmd SchemeReport) CommandClassID() byte {
	return 0x98
}

func (cmd SchemeReport) CommandID() byte {
	return byte(CommandSchemeReport)
}

func (cmd *SchemeReport) UnmarshalBinary(data []byte) error {
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

	cmd.SupportedSecuritySchemes = payload[i]
	i++

	return nil
}

func (cmd *SchemeReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = cmd.CommandClassID()
	payload[1] = cmd.CommandID()

	payload = append(payload, cmd.SupportedSecuritySchemes)

	return
}