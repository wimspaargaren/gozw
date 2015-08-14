// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package securitypanelzone

// <no value>

type SecurityPanelZoneSupportedReport struct {
	ZonesSupported byte

	Zm bool
}

func ParseSecurityPanelZoneSupportedReport(payload []byte) SecurityPanelZoneSupportedReport {
	val := SecurityPanelZoneSupportedReport{}

	i := 2

	val.ZonesSupported = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		val.Zm = true
	} else {
		val.Zm = false
	}

	i += 1

	return val
}