package application

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"github.com/wimspaargaren/gozw/cc"
	"github.com/wimspaargaren/gozw/cc/association"
	"github.com/wimspaargaren/gozw/cc/battery"
	"github.com/wimspaargaren/gozw/cc/manufacturer-specific"
	"github.com/wimspaargaren/gozw/cc/manufacturer-specific-v2"
	"github.com/wimspaargaren/gozw/cc/security"
	"github.com/wimspaargaren/gozw/cc/version"
	"github.com/wimspaargaren/gozw/cc/version-v2"
	"github.com/wimspaargaren/gozw/protocol"
	"github.com/wimspaargaren/gozw/serial-api"
	"github.com/wimspaargaren/gozw/util"
)

// Node is an in-memory representation of a Z-Wave node
type Node struct {
	NodeID byte

	Capability          byte
	BasicDeviceClass    byte
	GenericDeviceClass  byte
	SpecificDeviceClass byte

	Failing bool

	CommandClasses cc.CommandClassSet

	NetworkKeySent bool

	ManufacturerID uint16
	ProductTypeID  uint16
	ProductID      uint16

	QueryStageSecurity     bool
	QueryStageManufacturer bool
	QueryStageVersions     bool

	queryStageSecurityComplete     chan bool
	queryStageManufacturerComplete chan bool
	queryStageVersionsComplete     chan bool

	application *Layer
}

func NewNode(application *Layer, nodeID byte) (*Node, error) {
	node := &Node{
		NodeID: nodeID,

		CommandClasses: cc.CommandClassSet{},

		QueryStageSecurity:     false,
		QueryStageManufacturer: false,
		QueryStageVersions:     false,

		queryStageSecurityComplete:     make(chan bool),
		queryStageManufacturerComplete: make(chan bool),
		queryStageVersionsComplete:     make(chan bool),

		application: application,
	}

	err := node.loadFromDb()
	if err != nil {
		initErr := node.initialize()
		if initErr != nil {
			return nil, initErr
		}

		node.saveToDb()
	}

	return node, nil
}

func (n *Node) loadFromDb() error {
	var data []byte
	err := n.application.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		data = bucket.Get([]byte{n.NodeID})

		if len(data) == 0 {
			return errors.New("Node not found")
		}

		return nil
	})

	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(data, n)
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) initialize() error {
	nodeInfo, err := n.application.serialAPI.GetNodeProtocolInfo(n.NodeID)
	if err != nil {
		fmt.Println(err)
	} else {
		n.setFromNodeProtocolInfo(nodeInfo)
	}

	if n.NodeID == 1 {
		// self is never failing
		n.Failing = false
	} else {
		failing, err := n.application.serialAPI.IsFailedNode(n.NodeID)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		n.Failing = failing
	}

	return n.saveToDb()
}

func (n *Node) saveToDb() error {
	data, err := msgpack.Marshal(n)
	if err != nil {
		return err
	}

	return n.application.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("nodes"))
		return bucket.Put([]byte{n.NodeID}, data)
	})
}

func (n *Node) IsSecure() bool {
	return n.CommandClasses.Supports(cc.Security)
}

func (n *Node) IsListening() bool {
	return n.Capability&0x80 == 0x80
}

func (n *Node) GetBasicDeviceClassName() string {
	return protocol.GetBasicDeviceTypeName(n.BasicDeviceClass)
}

func (n *Node) GetGenericDeviceClassName() string {
	return protocol.GetGenericDeviceTypeName(n.GenericDeviceClass)
}

func (n *Node) GetSpecificDeviceClassName() string {
	return protocol.GetSpecificDeviceTypeName(n.GenericDeviceClass, n.SpecificDeviceClass)
}

func (n *Node) SendCommand(command cc.Command) error {
	commandClass := cc.CommandClassID(command.CommandClassID())

	if commandClass == cc.Security {
		switch command.(type) {
		case *security.CommandsSupportedGet, *security.CommandsSupportedReport:
			return n.application.SendDataSecure(n.NodeID, command)
		}
	}

	if !n.CommandClasses.Supports(commandClass) {
		return errors.New("Command class not supported")
	}

	if n.CommandClasses.IsSecure(commandClass) {
		return n.application.SendDataSecure(n.NodeID, command)
	}

	return n.application.SendData(n.NodeID, command)
}

func (n *Node) SendRawCommand(payload []byte) error {
	commandClass := cc.CommandClassID(payload[0])

	if !n.CommandClasses.Supports(commandClass) {
		return errors.New("Command class not supported")
	}

	if n.CommandClasses.IsSecure(commandClass) {
		return n.application.SendDataSecure(n.NodeID, util.ByteMarshaler(payload))
	}

	return n.application.SendData(n.NodeID, util.ByteMarshaler(payload))
}

func (n *Node) AddAssociation(groupID byte, nodeIDs ...byte) error {
	// sort of an arbitrary limit for now, but I'm not sure what it should be
	if len(nodeIDs) > 20 {
		return errors.New("Too many associated nodes")
	}

	fmt.Println("Associating")

	return n.SendCommand(&association.Set{
		GroupingIdentifier: groupID,
		NodeId:             nodeIDs,
	})
}

func (n *Node) LoadSupportedSecurityCommands() error {
	return n.application.SendDataSecure(n.NodeID, &security.CommandsSupportedGet{})
}

func (n *Node) RequestNodeInformationFrame() error {
	return n.application.serialAPI.RequestNodeInfo(n.NodeID)
}

func (n *Node) LoadCommandClassVersions() error {
	for _, commandClass := range n.CommandClasses {
		time.Sleep(1 * time.Second)
		cmd := &version.CommandClassGet{RequestedCommandClass: byte(commandClass.CommandClass)}
		var err error

		if !commandClass.Secure {
			err = n.application.SendData(n.NodeID, cmd)
		} else {
			err = n.application.SendDataSecure(n.NodeID, cmd)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Node) LoadManufacturerInfo() error {
	return n.SendCommand(&manufacturerspecific.Get{})
}

func (n *Node) nextQueryStage() {
	if !n.QueryStageSecurity && n.IsSecure() {
		n.LoadSupportedSecurityCommands()
		return
	}

	if !n.QueryStageVersions {
		n.LoadCommandClassVersions()
		return
	}

	if !n.QueryStageManufacturer {
		n.LoadManufacturerInfo()
		return
	}
}

func (n *Node) emitNodeEvent(event cc.Command) {
	buf, err := event.MarshalBinary()
	if err != nil {
		fmt.Printf("error encoding: %v\n", err)
		return
	}

	n.application.EventBus.Publish("nodeCommand", n.NodeID, buf)
}

func (n *Node) receiveControllerUpdate(update serialapi.ControllerUpdate) {
	n.setFromApplicationControllerUpdate(update)
	n.saveToDb()
}

func (n *Node) setFromAddNodeCallback(nodeInfo *serialapi.AddRemoveNodeCallback) {
	n.NodeID = nodeInfo.Source
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cmd := range nodeInfo.CommandClasses {
		n.CommandClasses.Add(cc.CommandClassID(cmd))
	}

	n.saveToDb()
}

func (n *Node) setFromApplicationControllerUpdate(nodeInfo serialapi.ControllerUpdate) {
	n.BasicDeviceClass = nodeInfo.Basic
	n.GenericDeviceClass = nodeInfo.Generic
	n.SpecificDeviceClass = nodeInfo.Specific

	for _, cmd := range nodeInfo.CommandClasses {
		n.CommandClasses.Add(cc.CommandClassID(cmd))
	}

	n.saveToDb()
}

func (n *Node) setFromNodeProtocolInfo(nodeInfo *serialapi.NodeProtocolInfo) {
	n.Capability = nodeInfo.Capability
	n.BasicDeviceClass = nodeInfo.BasicDeviceClass
	n.GenericDeviceClass = nodeInfo.GenericDeviceClass
	n.SpecificDeviceClass = nodeInfo.SpecificDeviceClass

	n.saveToDb()
}

func (n *Node) receiveSecurityCommandsSupportedReport(cmd security.CommandsSupportedReport) {
	for _, supported := range cmd.CommandClassSupport {
		n.CommandClasses.SetSecure(cc.CommandClassID(supported), true)
	}

	select {
	case n.queryStageSecurityComplete <- true:
	default:
	}

	n.QueryStageSecurity = true
	n.saveToDb()
	n.nextQueryStage()
}

func (n *Node) receiveManufacturerInfo(mfgId, productTypeId, productId uint16) {
	n.ManufacturerID = mfgId
	n.ProductTypeID = productTypeId
	n.ProductID = productId

	select {
	case n.queryStageManufacturerComplete <- true:
	default:
	}

	n.QueryStageManufacturer = true
	n.saveToDb()
	n.nextQueryStage()
}

func (n *Node) receiveCommandClassVersion(id cc.CommandClassID, version uint8) {
	n.CommandClasses.SetVersion(id, version)

	if n.CommandClasses.AllVersionsReceived() {
		select {
		case n.queryStageVersionsComplete <- true:
		default:
		}

		n.QueryStageVersions = true
		defer n.nextQueryStage()
	}

	n.saveToDb()
}

func (n *Node) receiveApplicationCommand(cmd serialapi.ApplicationCommand) {
	commandClassID := cc.CommandClassID(cmd.CommandData[0])
	ver := n.CommandClasses.GetVersion(commandClassID)
	if ver == 0 {
		ver = 1

		if !(commandClassID == cc.Version || commandClassID == cc.Security) {
			fmt.Printf("error: no version loaded for %s\n", commandClassID)
		}
	}

	command, err := cc.Parse(ver, cmd.CommandData)
	if err != nil {
		fmt.Println("error parsing command class", err)
		return
	}

	switch command.(type) {

	case *battery.Report:
		if cmd.CommandData[2] == 0xFF {
			fmt.Printf("Node %d: low battery alert\n", n.NodeID)
		} else {
			fmt.Printf("Node %d: battery level is %d\n", n.NodeID, command.(*battery.Report))
		}
		n.emitNodeEvent(command)

	case *security.CommandsSupportedReport:
		fmt.Println("security commands supported report")
		n.receiveSecurityCommandsSupportedReport(*command.(*security.CommandsSupportedReport))
		fmt.Println(n.GetSupportedSecureCommandClassStrings())

	case *manufacturerspecific.Report:
		spew.Dump(command.(*manufacturerspecific.Report))
		report := *command.(*manufacturerspecific.Report)
		n.receiveManufacturerInfo(report.ManufacturerId, report.ProductTypeId, report.ProductId)
		n.emitNodeEvent(command)

	case *manufacturerspecificv2.Report:
		spew.Dump(command.(*manufacturerspecificv2.Report))
		report := *command.(*manufacturerspecificv2.Report)
		n.receiveManufacturerInfo(report.ManufacturerId, report.ProductTypeId, report.ProductId)
		n.emitNodeEvent(command)

	case *version.CommandClassReport:
		spew.Dump(command.(*version.CommandClassReport))
		report := command.(*version.CommandClassReport)
		n.receiveCommandClassVersion(cc.CommandClassID(report.RequestedCommandClass), report.CommandClassVersion)
		n.saveToDb()

	case *versionv2.CommandClassReport:
		spew.Dump(command.(*versionv2.CommandClassReport))
		report := command.(*versionv2.CommandClassReport)
		n.receiveCommandClassVersion(cc.CommandClassID(report.RequestedCommandClass), report.CommandClassVersion)
		n.saveToDb()

		// case alarm.Report:
		// 	spew.Dump(command.(alarm.Report))
		//
		// case usercode.Report:
		// 	spew.Dump(command.(usercode.Report))
		//
		// case doorlock.OperationReport:
		// 	spew.Dump(command.(doorlock.OperationReport))
		//
		// case thermostatmode.Report:
		// 	spew.Dump(command.(thermostatmode.Report))
		//
		// case thermostatoperatingstate.Report:
		// 	spew.Dump(command.(thermostatoperatingstate.Report))
		//
		// case thermostatsetpoint.Report:
		// 	spew.Dump(command.(thermostatsetpoint.Report))

	default:
		spew.Dump(command)
		n.emitNodeEvent(command)
	}
}

func (n *Node) String() string {
	str := fmt.Sprintf("Node %d: \n", n.NodeID)
	str += fmt.Sprintf("  Failing? %t\n", n.Failing)
	str += fmt.Sprintf("  Is listening? %t\n", n.IsListening())
	str += fmt.Sprintf("  Is secure? %t\n", n.IsSecure())
	str += fmt.Sprintf("  Basic device class: %s\n", n.GetBasicDeviceClassName())
	str += fmt.Sprintf("  Generic device class: %s\n", n.GetGenericDeviceClassName())
	str += fmt.Sprintf("  Specific device class: %s\n", n.GetSpecificDeviceClassName())
	str += fmt.Sprintf("  Manufacturer ID: %#x\n", n.ManufacturerID)
	str += fmt.Sprintf("  Product Type ID: %#x\n", n.ProductTypeID)
	str += fmt.Sprintf("  Product ID: %#x\n", n.ProductID)
	str += fmt.Sprintf("  Supported command classes:\n")

	for _, cmd := range n.CommandClasses {
		if cmd.Secure {
			str += fmt.Sprintf("    - %s (v%d) (secure)\n", cmd.CommandClass.String(), cmd.Version)
		} else {
			str += fmt.Sprintf("    - %s (v%d)\n", cmd.CommandClass.String(), cmd.Version)
		}
	}

	return str
}

func (n *Node) GetSupportedCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.CommandClasses.ListBySecureStatus(false))
	if len(strings) == 0 {
		return []string{
			"None (probably not loaded; need to request a NIF)",
		}
	}

	return strings
}

func (n *Node) GetSupportedSecureCommandClassStrings() []string {
	strings := commandClassSetToStrings(n.CommandClasses.ListBySecureStatus(true))
	return strings
}

func commandClassSetToStrings(commandClasses []cc.CommandClassID) []string {
	if len(commandClasses) == 0 {
		return []string{}
	}

	ccStrings := []string{}

	for _, cmd := range commandClasses {
		ccStrings = append(ccStrings, cmd.String())
	}

	return ccStrings
}
