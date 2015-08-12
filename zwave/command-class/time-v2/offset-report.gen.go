// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package timev2

// <no value>

type TimeOffsetReport struct {
	HourTzo byte

	SignTzo bool

	MinuteTzo byte

	MinuteOffsetDst byte

	SignOffsetDst bool

	MonthStartDst byte

	DayStartDst byte

	HourStartDst byte

	MonthEndDst byte

	DayEndDst byte

	HourEndDst byte
}

func ParseTimeOffsetReport(payload []byte) TimeOffsetReport {
	val := TimeOffsetReport{}

	i := 2

	val.HourTzo = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.SignTzo = true
	} else {
		val.SignTzo = false
	}

	i += 1

	val.MinuteTzo = payload[i]
	i++

	val.MinuteOffsetDst = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.SignOffsetDst = true
	} else {
		val.SignOffsetDst = false
	}

	i += 1

	val.MonthStartDst = payload[i]
	i++

	val.DayStartDst = payload[i]
	i++

	val.HourStartDst = payload[i]
	i++

	val.MonthEndDst = payload[i]
	i++

	val.DayEndDst = payload[i]
	i++

	val.HourEndDst = payload[i]
	i++

	return val
}
