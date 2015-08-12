// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatsetpointv2

// <no value>

type ThermostatSetpointReport struct {
	SetpointType byte

	Size byte

	Scale byte

	Precision byte

	Value []byte
}

func ParseThermostatSetpointReport(payload []byte) ThermostatSetpointReport {
	val := ThermostatSetpointReport{}

	i := 2

	val.SetpointType = (payload[i] & 0x0F)

	i += 1

	val.Size = (payload[i] & 0x07)

	val.Scale = (payload[i] & 0x18) << 3

	val.Precision = (payload[i] & 0xE0) << 5

	i += 1

	val.Value = payload[i : i+1]
	i += 1

	return val
}
