// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationv2

import "encoding/gob"

func init() {
	gob.Register(AssociationSpecificGroupGet{})
}

// <no value>
type AssociationSpecificGroupGet struct {
}

func (cmd *AssociationSpecificGroupGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *AssociationSpecificGroupGet) MarshalBinary() (payload []byte, err error) {

	return
}
