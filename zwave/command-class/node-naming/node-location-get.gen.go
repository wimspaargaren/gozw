// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package nodenaming

import "encoding/gob"

func init() {
	gob.Register(NodeNamingNodeLocationGet{})
}

// <no value>
type NodeNamingNodeLocationGet struct {
}

func (cmd *NodeNamingNodeLocationGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *NodeNamingNodeLocationGet) MarshalBinary() (payload []byte, err error) {

	return
}
