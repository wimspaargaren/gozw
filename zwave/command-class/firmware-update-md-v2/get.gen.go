// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package firmwareupdatemdv2

// <no value>

type FirmwareUpdateMdGet struct {
	NumberOfReports byte

	ReportNumber1 byte

	Zero bool

	ReportNumber2 byte
}

func ParseFirmwareUpdateMdGet(payload []byte) FirmwareUpdateMdGet {
	val := FirmwareUpdateMdGet{}

	i := 2

	val.NumberOfReports = payload[i]
	i++

	val.ReportNumber1 = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Zero = true
	} else {
		val.Zero = false
	}

	i += 1

	val.ReportNumber2 = payload[i]
	i++

	return val
}
