// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package applicationstatus

// <no value>

type ApplicationRejectedRequest struct {
	Status byte
}

func ParseApplicationRejectedRequest(payload []byte) ApplicationRejectedRequest {
	val := ApplicationRejectedRequest{}

	i := 2

	val.Status = payload[i]
	i++

	return val
}
