// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchmultilevel

// <no value>

type SwitchMultilevelReport struct {
	Value byte
}

func ParseSwitchMultilevelReport(payload []byte) SwitchMultilevelReport {
	val := SwitchMultilevelReport{}

	i := 2

	val.Value = payload[i]
	i++

	return val
}
