// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

// <no value>

type ChimneyFanSpeedSet struct {
	Speed byte
}

func ParseChimneyFanSpeedSet(payload []byte) ChimneyFanSpeedSet {
	val := ChimneyFanSpeedSet{}

	i := 2

	val.Speed = payload[i]
	i++

	return val
}
