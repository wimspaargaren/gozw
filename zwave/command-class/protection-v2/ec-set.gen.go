// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package protectionv2

// <no value>

type ProtectionEcSet struct {
	NodeId byte
}

func ParseProtectionEcSet(payload []byte) ProtectionEcSet {
	val := ProtectionEcSet{}

	i := 2

	val.NodeId = payload[i]
	i++

	return val
}
