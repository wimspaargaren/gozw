// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scenecontrollerconf

import (
	"encoding/gob"
	"errors"
)

func init() {
	gob.Register(SceneControllerConfSet{})
}

// <no value>
type SceneControllerConfSet struct {
	GroupId byte

	SceneId byte

	DimmingDuration byte
}

func (cmd *SceneControllerConfSet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)
	i := 0

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SceneId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DimmingDuration = payload[i]
	i++

	return nil
}

func (cmd *SceneControllerConfSet) MarshalBinary() (payload []byte, err error) {

	payload = append(payload, cmd.GroupId)

	payload = append(payload, cmd.SceneId)

	payload = append(payload, cmd.DimmingDuration)

	return
}
