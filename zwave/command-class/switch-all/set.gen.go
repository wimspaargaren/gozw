// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchall

// <no value>

type SwitchAllSet struct {
	Mode byte
}

func ParseSwitchAllSet(payload []byte) SwitchAllSet {
	val := SwitchAllSet{}

	i := 2

	val.Mode = payload[i]
	i++

	return val
}
