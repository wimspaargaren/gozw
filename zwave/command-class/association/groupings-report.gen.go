// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package association

// <no value>

type AssociationGroupingsReport struct {
	SupportedGroupings byte
}

func ParseAssociationGroupingsReport(payload []byte) AssociationGroupingsReport {
	val := AssociationGroupingsReport{}

	i := 2

	val.SupportedGroupings = payload[i]
	i++

	return val
}
