// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sceneactuatorconf

// <no value>

type SceneActuatorConfGet struct {
	SceneId byte
}

func ParseSceneActuatorConfGet(payload []byte) SceneActuatorConfGet {
	val := SceneActuatorConfGet{}

	i := 2

	val.SceneId = payload[i]
	i++

	return val
}
