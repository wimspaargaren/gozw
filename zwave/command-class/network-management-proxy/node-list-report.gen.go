// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementproxy

import "errors"

// <no value>

type NodeListReport struct {
	SeqNo byte

	Status byte

	NodeListControllerId byte

	NodeListData byte
}

func (cmd *NodeListReport) UnmarshalBinary(payload []byte) error {
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Status = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeListControllerId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeListData = payload[i]
	i++

	return nil
}

func (cmd *NodeListReport) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Status)

	payload = append(payload, cmd.NodeListControllerId)

	payload = append(payload, cmd.NodeListData)

	return
}