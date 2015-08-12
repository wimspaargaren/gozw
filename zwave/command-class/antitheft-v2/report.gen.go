// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package antitheftv2

import "encoding/binary"

// <no value>

type AntitheftReport struct {
	AntiTheftProtectionStatus byte

	ManufacturerId uint16

	AntiTheftHintNumberBytes byte

	AntiTheftHintByte []byte
}

func ParseAntitheftReport(payload []byte) AntitheftReport {
	val := AntitheftReport{}

	i := 2

	val.AntiTheftProtectionStatus = payload[i]
	i++

	val.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.AntiTheftHintNumberBytes = payload[i]
	i++

	val.AntiTheftHintByte = payload[i : i+2]
	i += 2

	return val
}
