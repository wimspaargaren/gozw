// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package protectionv2

// <no value>

type ProtectionReport struct {
	LocalProtectionState byte

	RfProtectionState byte
}

func ParseProtectionReport(payload []byte) ProtectionReport {
	val := ProtectionReport{}

	i := 2

	val.LocalProtectionState = (payload[i] & 0x0F)

	i += 1

	val.RfProtectionState = (payload[i] & 0x0F)

	i += 1

	return val
}
