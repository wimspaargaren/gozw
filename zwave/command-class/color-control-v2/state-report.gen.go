// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrolv2

// <no value>

type StateReport struct {
	CapabilityId byte

	State byte
}

func ParseStateReport(payload []byte) StateReport {
	val := StateReport{}

	i := 2

	val.CapabilityId = payload[i]
	i++

	val.State = payload[i]
	i++

	return val
}
