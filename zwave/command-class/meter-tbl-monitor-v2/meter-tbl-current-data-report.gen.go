// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblmonitorv2

import "encoding/binary"

// <no value>

type MeterTblCurrentDataReport struct {
	ReportsToFollow byte

	RateType byte

	OperatingStatusIndication bool

	Dataset uint32

	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte
}

func ParseMeterTblCurrentDataReport(payload []byte) MeterTblCurrentDataReport {
	val := MeterTblCurrentDataReport{}

	i := 2

	val.ReportsToFollow = payload[i]
	i++

	val.RateType = (payload[i] & 0x03)

	if payload[i]&0x80 == 0x80 {
		val.OperatingStatusIndication = true
	} else {
		val.OperatingStatusIndication = false
	}

	i += 1

	val.Dataset = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

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

	return val
}