// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementinclusion

// <no value>

type ReturnRouteAssignComplete struct {
	SeqNo byte

	Status byte
}

func ParseReturnRouteAssignComplete(payload []byte) ReturnRouteAssignComplete {
	val := ReturnRouteAssignComplete{}

	i := 2

	val.SeqNo = payload[i]
	i++

	val.Status = payload[i]
	i++

	return val
}
