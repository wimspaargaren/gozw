// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrolv2

// <no value>

type StartCapabilityLevelChange struct {
	Res1 byte

	IgnoreStartState bool

	Updown bool

	Res2 bool

	CapabilityId byte

	StartState byte
}

func ParseStartCapabilityLevelChange(payload []byte) StartCapabilityLevelChange {
	val := StartCapabilityLevelChange{}

	i := 2

	val.Res1 = (payload[i] & 0x1F)

	if payload[i]&0x20 == 0x20 {
		val.IgnoreStartState = true
	} else {
		val.IgnoreStartState = false
	}

	if payload[i]&0x40 == 0x40 {
		val.Updown = true
	} else {
		val.Updown = false
	}

	if payload[i]&0x80 == 0x80 {
		val.Res2 = true
	} else {
		val.Res2 = false
	}

	i += 1

	val.CapabilityId = payload[i]
	i++

	val.StartState = payload[i]
	i++

	return val
}
