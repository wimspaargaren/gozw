// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv3

// <no value>

type MultiChannelEndPointReport struct {
	Identical bool

	Dynamic bool

	EndPoints byte
}

func ParseMultiChannelEndPointReport(payload []byte) MultiChannelEndPointReport {
	val := MultiChannelEndPointReport{}

	i := 2

	if payload[i]&0x40 == 0x40 {
		val.Identical = true
	} else {
		val.Identical = false
	}

	if payload[i]&0x80 == 0x80 {
		val.Dynamic = true
	} else {
		val.Dynamic = false
	}

	i += 1

	val.EndPoints = (payload[i] & 0x7F)

	i += 1

	return val
}
