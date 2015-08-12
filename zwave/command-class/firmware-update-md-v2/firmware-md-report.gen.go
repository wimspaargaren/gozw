// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package firmwareupdatemdv2

import "encoding/binary"

// <no value>

type FirmwareMdReport struct {
	ManufacturerId uint16

	FirmwareId uint16

	Checksum uint16
}

func ParseFirmwareMdReport(payload []byte) FirmwareMdReport {
	val := FirmwareMdReport{}

	i := 2

	val.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.FirmwareId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.Checksum = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return val
}
