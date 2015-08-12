// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeup

import "encoding/binary"

// <no value>

type WakeUpIntervalReport struct {
	Seconds uint32

	Nodeid byte
}

func ParseWakeUpIntervalReport(payload []byte) WakeUpIntervalReport {
	val := WakeUpIntervalReport{}

	i := 2

	val.Seconds = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	val.Nodeid = payload[i]
	i++

	return val
}
