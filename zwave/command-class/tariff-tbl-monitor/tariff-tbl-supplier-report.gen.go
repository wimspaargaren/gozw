// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package tarifftblmonitor

import "encoding/binary"

// <no value>

type TariffTblSupplierReport struct {
	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte

	Currency uint32

	StandingChargePeriod byte

	StandingChargePrecision byte

	StandingChargeValue uint32

	NumberOfSupplierCharacters byte

	SupplierCharacter []byte
}

func ParseTariffTblSupplierReport(payload []byte) TariffTblSupplierReport {
	val := TariffTblSupplierReport{}

	i := 2

	val.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.Month = payload[i]
	i++

	val.Day = payload[i]
	i++

	val.HourLocalTime = payload[i]
	i++

	val.MinuteLocalTime = payload[i]
	i++

	val.SecondLocalTime = payload[i]
	i++

	val.Currency = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	val.StandingChargePeriod = (payload[i] & 0x1F)

	val.StandingChargePrecision = (payload[i] & 0xE0) << 5

	i += 1

	val.StandingChargeValue = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	val.NumberOfSupplierCharacters = (payload[i] & 0x1F)

	i += 1

	val.SupplierCharacter = payload[i : i+9]
	i += 9

	return val
}
