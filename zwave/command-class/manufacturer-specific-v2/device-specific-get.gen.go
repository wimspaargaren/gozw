// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package manufacturerspecificv2

// <no value>

type DeviceSpecificGet struct {
	DeviceIdType byte
}

func ParseDeviceSpecificGet(payload []byte) DeviceSpecificGet {
	val := DeviceSpecificGet{}

	i := 2

	val.DeviceIdType = (payload[i] & 0x07)

	i += 1

	return val
}