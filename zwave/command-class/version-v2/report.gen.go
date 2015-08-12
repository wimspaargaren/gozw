// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package versionv2

// <no value>

type VersionReport struct {
	ZWaveLibraryType byte

	ZWaveProtocolVersion byte

	ZWaveProtocolSubVersion byte

	Firmware0Version byte

	Firmware0SubVersion byte

	HardwareVersion byte

	NumberOfFirmwareTargets byte
}

func ParseVersionReport(payload []byte) VersionReport {
	val := VersionReport{}

	i := 2

	val.ZWaveLibraryType = payload[i]
	i++

	val.ZWaveProtocolVersion = payload[i]
	i++

	val.ZWaveProtocolSubVersion = payload[i]
	i++

	val.Firmware0Version = payload[i]
	i++

	val.Firmware0SubVersion = payload[i]
	i++

	val.HardwareVersion = payload[i]
	i++

	val.NumberOfFirmwareTargets = payload[i]
	i++

	return val
}
