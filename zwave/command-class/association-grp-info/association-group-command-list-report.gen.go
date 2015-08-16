// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationgrpinfo

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(AssociationGroupCommandListReport{})
}

// <no value>
type AssociationGroupCommandListReport struct {
	GroupingIdentifier byte

	ListLength byte

	Command []byte
}

func (cmd *AssociationGroupCommandListReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ListLength = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Command = payload[i : i+1]
	i += 1

	return nil
}

func (cmd *AssociationGroupCommandListReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.ListLength)

	if cmd.Command != nil && len(cmd.Command) > 0 {
		payload = append(payload, cmd.Command...)
	}

	return
}
