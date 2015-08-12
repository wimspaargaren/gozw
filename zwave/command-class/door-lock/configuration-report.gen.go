// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package doorlock

// <no value>

type DoorLockConfigurationReport struct {
	OperationType byte

	InsideDoorHandlesState byte

	OutsideDoorHandlesState byte

	LockTimeoutMinutes byte

	LockTimeoutSeconds byte
}

func ParseDoorLockConfigurationReport(payload []byte) DoorLockConfigurationReport {
	val := DoorLockConfigurationReport{}

	i := 2

	val.OperationType = payload[i]
	i++

	val.InsideDoorHandlesState = (payload[i] & 0x0F)

	val.OutsideDoorHandlesState = (payload[i] & 0xF0) << 4

	i += 1

	val.LockTimeoutMinutes = payload[i]
	i++

	val.LockTimeoutSeconds = payload[i]
	i++

	return val
}
