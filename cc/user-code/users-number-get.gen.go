// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package usercode

import (
	"encoding/gob"

	"github.com/wimspaargaren/gozw/cc"
)

const CommandUsersNumberGet cc.CommandID = 0x04

func init() {
	gob.Register(UsersNumberGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x63),
		Command:      cc.CommandID(0x04),
		Version:      1,
	}, NewUsersNumberGet)
}

func NewUsersNumberGet() cc.Command {
	return &UsersNumberGet{}
}

// <no value>
type UsersNumberGet struct {
}

func (cmd UsersNumberGet) CommandClassID() cc.CommandClassID {
	return 0x63
}

func (cmd UsersNumberGet) CommandID() cc.CommandID {
	return CommandUsersNumberGet
}

func (cmd UsersNumberGet) CommandIDString() string {
	return "USERS_NUMBER_GET"
}

func (cmd *UsersNumberGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *UsersNumberGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
