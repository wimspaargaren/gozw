// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package dcpmonitor

import "encoding/binary"

// <no value>

type DcpEventStatusReport struct {
	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte

	EventStatus byte
}

func ParseDcpEventStatusReport(payload []byte) DcpEventStatusReport {
	val := DcpEventStatusReport{}

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

	val.EventStatus = payload[i]
	i++

	return val
}
