// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package avcontentdirectorymd

import "encoding/gob"

func init() {
	gob.Register(AvContentBrowseMdByLetterReport{})
}

// <no value>
type AvContentBrowseMdByLetterReport struct {
}

func (cmd *AvContentBrowseMdByLetterReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning
	payload := make([]byte, len(data))
	copy(payload, data)

	return nil
}

func (cmd *AvContentBrowseMdByLetterReport) MarshalBinary() (payload []byte, err error) {

	return
}
