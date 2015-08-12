// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv3

// <no value>

type MultiChannelCmdEncap struct {
	SourceEndPoint byte

	DestinationEndPoint byte

	BitAddress bool

	CommandClass byte

	Command byte

	Parameter []byte
}

func ParseMultiChannelCmdEncap(payload []byte) MultiChannelCmdEncap {
	val := MultiChannelCmdEncap{}

	i := 2

	val.SourceEndPoint = (payload[i] & 0x7F)

	i += 1

	val.DestinationEndPoint = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.BitAddress = true
	} else {
		val.BitAddress = false
	}

	i += 1

	val.CommandClass = payload[i]
	i++

	val.Command = payload[i]
	i++

	val.Parameter = payload[i:]

	return val
}
