// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatsetpointv2

// <no value>

type ThermostatSetpointSupportedReport struct {
	BitMask byte
}

func ParseThermostatSetpointSupportedReport(payload []byte) ThermostatSetpointSupportedReport {
	val := ThermostatSetpointSupportedReport{}

	i := 2

	val.BitMask = payload[i]
	i++

	return val
}
