// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatheating

// <no value>

type ThermostatHeatingStatusReport struct {
	Status byte
}

func ParseThermostatHeatingStatusReport(payload []byte) ThermostatHeatingStatusReport {
	val := ThermostatHeatingStatusReport{}

	i := 2

	val.Status = payload[i]
	i++

	return val
}
