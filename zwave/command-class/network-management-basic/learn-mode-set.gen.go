// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package networkmanagementbasic

// <no value>

type LearnModeSet struct {
	SeqNo byte

	Mode byte
}

func ParseLearnModeSet(payload []byte) LearnModeSet {
	val := LearnModeSet{}

	i := 2

	val.SeqNo = payload[i]
	i++

	val.Mode = payload[i]
	i++

	return val
}
