package vda5050

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/1000in1/m/logger"
	mqttclient "github.com/1000in1/m/mqtt_client"
	"github.com/1000in1/m/vda5050/connection"
	"github.com/1000in1/m/vda5050/instantactions"
	"github.com/1000in1/m/vda5050/order"
	"github.com/1000in1/m/vda5050/state"
)

type VDA5050 struct {
	tag           string
	State         state.State
	TaskId        string
	FinalMovement bool
	FinalOrder    bool
	DeviceId      string
	logger        logger.LoggerIF
	headerID      int64
	lock          sync.Mutex
	TopicPrefix   string
	client        *mqttclient.MqttClient
}

func NewVDA5050(deviceID string, manufacturer string, ver string) *VDA5050 {

	return &VDA5050{
		State: state.State{
			ActionStates: []*state.ActionState{},
			Errors:       []state.Error{},
			EdgeStates:   []state.EdgeState{},
			NodeStates:   []*state.NodeState{},
			Loads:        []state.Load{},
			Information:  []state.Information{},
			Version:      ver,
			Manufacturer: manufacturer,
			SerialNumber: deviceID,
		},
		tag:         "VDA5050",
		TaskId:      "",
		DeviceId:    deviceID,
		logger:      nil,
		headerID:    0,
		TopicPrefix: "uagv/v2/" + manufacturer + "/" + deviceID,
		client:      nil,
	}
}

func (v *VDA5050) GetNextHeaderId() int64 {
	// 使用 atomic.AddInt64 来安全地递增 headerId
	return atomic.AddInt64(&v.headerID, 1)
}

func (v *VDA5050) INFO(message string) {
	if v.logger != nil {
		v.logger.INFO(v.tag, message)
	}
}

func (v *VDA5050) ERROR(message string) {
	if v.logger != nil {
		v.logger.ERROR(v.tag, message)
	}
}

func (v *VDA5050) SetLogger(l *logger.Logger) {
	v.logger = l
}

func (v *VDA5050) UpdateInstantActions(insAction *instantactions.InstantActions) {

	v.lock.Lock()
	defer v.lock.Unlock()

	for _, action := range insAction.Actions {

		ac := state.ActionState{
			ActionID:         *action.ActionID,
			ActionType:       *action.ActionType,
			BlockingType:     string(*action.BlockingType),
			ActionStatus:     state.Waiting,
			ActionParameters: action.ActionParameters,
		}

		v.State.ActionStates = append(v.State.ActionStates, &ac)
	}

}

func (v *VDA5050) UpdateOrder(vda5050_order *order.Order) {

	v.lock.Lock()
	defer v.lock.Unlock()

	if v.State.OrderID == *vda5050_order.OrderID && v.State.OrderUpdateID == *vda5050_order.OrderUpdateID {
		return
	}

	v.State.OrderID = *vda5050_order.OrderID
	v.State.OrderUpdateID = *vda5050_order.OrderUpdateID

	////Custom properties
	if vda5050_order.TaskId != nil {
		v.TaskId = *vda5050_order.TaskId
	}

	if vda5050_order.FinalMovement != nil {
		v.FinalMovement = *vda5050_order.FinalMovement
	}
	if vda5050_order.FinalOrder != nil {
		v.FinalOrder = *vda5050_order.FinalOrder
	}

	// Clear action states when order changes.
	v.State.ActionStates = []*state.ActionState{}

	for _, node := range vda5050_order.Nodes {

		nodeState := &state.NodeState{
			NodeID:     *node.NodeID,
			SequenceID: *node.SequenceID,
			Released:   *node.Released,
		}

		nodeState.NodePosition.X = *node.NodePosition.X
		nodeState.NodePosition.Y = *node.NodePosition.Y
		nodeState.NodePosition.MapID = *node.NodePosition.MapID

		length := len(v.State.NodeStates)
		if length > 0 {

			if nodeState.SequenceID == v.State.NodeStates[length-1].SequenceID && nodeState.NodeID == v.State.NodeStates[length-1].NodeID {
				continue
			}

		}

		v.State.NodeStates = append(v.State.NodeStates, nodeState)

		for _, action := range node.Actions {

			actionState := state.ActionState{
				ActionID:         *action.ActionID,
				ActionType:       *action.ActionType,
				BlockingType:     string(*action.BlockingType),
				ActionStatus:     state.Initializing,
				NodeID:           *node.NodeID,
				ActionParameters: action.ActionParameters,
			}

			v.State.ActionStates = append(v.State.ActionStates, &actionState)

		}

	}

	//v.INFO(fmt.Sprintf("NodeCount:%d", len(v.State.NodeStates)))

}

func (v *VDA5050) GetNodes() []*state.NodeState {
	return v.State.NodeStates
}

func (v *VDA5050) GetNodesCount() int { return len(v.State.NodeStates) }

func (v *VDA5050) GetTargetNode() *state.NodeState {

	if v.GetNodesCount() > 0 {
		return v.State.NodeStates[0]
	}

	return nil
}

func (v *VDA5050) PopNodeStates() *state.NodeState {

	if v.GetNodesCount() > 0 {
		target := v.GetTargetNode()
		v.State.NodeStates = v.State.NodeStates[1:]
		return target
	}
	return nil

}

func (v *VDA5050) GetActions() []*state.ActionState {
	return v.State.ActionStates
}

func (v *VDA5050) GetStateJsonString() (string, error) {

	v.State.HeaderID = v.GetNextHeaderId()

	json, err := v.State.Marshal()

	return string(json), err

}

func (v *VDA5050) GetConnectionJsonString(s connection.ConnectionState) ([]byte, error) {

	connection := connection.Connection{
		ConnectionState: s,
		HeaderID:        v.GetNextHeaderId(),
		Manufacturer:    v.State.Manufacturer,
		SerialNumber:    v.DeviceId,
		Version:         v.State.Version,
	}

	json, err := connection.Marshal()

	return json, err

}

func (v *VDA5050) GetBatteryState() *state.BatteryState { return &v.State.BatteryState }

func (v *VDA5050) ChangeBatteryState(charging bool, batteryHealth, batteryCharge, batteryVoltage float64) {

	v.State.BatteryState.Charging = charging
	v.State.BatteryState.BatteryHealth = batteryHealth
	v.State.BatteryState.BatteryCharge = batteryCharge
	v.State.BatteryState.BatteryVoltage = batteryVoltage
	v.State.BatteryState.Reach = 9999999

}

func (v *VDA5050) GetAgvPosition() *state.AgvPosition { return &v.State.AgvPosition }

func (v *VDA5050) GetSafetyState() *state.SafetyState { return &v.State.SafetyState }

func (v *VDA5050) SetSafetyState(estop state.EStop) {
	v.State.SafetyState.FieldViolation = true
	v.State.SafetyState.EStop = estop
}

func (v *VDA5050) SetOperatingMode(mode state.OperatingMode) {
	v.State.OperatingMode = mode
}

func (v *VDA5050) UpdateLastNode(lastNodeId string, lastNodeSequenceID *int64) {
	v.State.LastNodeID = lastNodeId
	if lastNodeSequenceID != nil {
		v.State.LastNodeSequenceID = *lastNodeSequenceID
	}
}

func (v *VDA5050) ChangeActionStatus(actionId string, status state.ActionStatus, cmdID string) bool {
	v.lock.Lock()
	defer v.lock.Unlock()

	for _, action := range v.State.ActionStates {
		if action.ActionID == actionId {
			//last := action.ActionStatus
			action.ActionStatus = status
			action.CmdID = cmdID
			return true
		}
	}

	return false
}

func (v *VDA5050) ClearActionStatus() {

	v.lock.Lock()
	defer v.lock.Unlock()

	temp := []*state.ActionState{}

	for _, action := range v.State.ActionStates {
		if action.ActionStatus == state.Finished || action.ActionStatus == state.Failed {
			continue
		}
		temp = append(temp, action)
	}

	v.State.ActionStates = temp

}

func (v *VDA5050) ChangeAgvPosition(mapId *string, x, y float64, theta *float64, initFlag bool) {
	v.lock.Lock()
	defer v.lock.Unlock()

	agv := v.GetAgvPosition()
	agv.X = x
	agv.Y = y

	if theta != nil {
		agv.Theta = *theta
	}
	if mapId != nil {
		agv.MapID = *mapId
	}
	agv.PositionInitialized = initFlag

}

func (v *VDA5050) ClearOrder() {
	v.lock.Lock()
	defer v.lock.Unlock()

	v.State.ActionStates = []*state.ActionState{}
	v.State.NodeStates = []*state.NodeState{}

}

func (v *VDA5050) AddLoad(loadID string, position string) {

	load := state.Load{
		LoadID:       loadID,
		LoadPosition: position,
	}

	v.State.Loads = append(v.State.Loads, load)

}

func (v *VDA5050) ClearLoad() {
	v.State.Loads = []state.Load{}
}

func (v *VDA5050) OnConnected(client mqttclient.MqttClient) {
	v.INFO("mqtt connected")

	client.Subscribe(v.TopicPrefix+"/order", 1)
	client.Subscribe(v.TopicPrefix+"/instantActions", 1)

	loginInfo, err := v.GetConnectionJsonString(connection.Online)
	if err != nil {
		v.ERROR("GetConnectionJsonString error:" + err.Error())
	} else {
		client.PublishEx(v.TopicPrefix+"/connection", loginInfo, byte(1), true)
	}

	v.client = &client

}

func (v *VDA5050) PublishState() {

	str, err := v.GetStateJsonString()
	if err == nil && v.client != nil {
		v.client.Publish(v.TopicPrefix+"/state", []byte(str))
	}

}

func (v *VDA5050) OnMessage(topic string, pyload []byte) {
	fmt.Println("recv:", topic, string(pyload))

	if topic == v.TopicPrefix+"/order" {
		vda5050_order, err := order.UnmarshalOrder(pyload)
		if err != nil {
			fmt.Println(err)
		} else {
			v.UpdateOrder(&vda5050_order)
			v.PublishState()
		}
	}

	if topic == v.TopicPrefix+"/instantActions" {

		vda5050_instantactions, err := instantactions.UnmarshalInstantActions(pyload)

		if err != nil {
			fmt.Println(err)
		} else {

			v.UpdateInstantActions(&vda5050_instantactions)
			v.PublishState()

		}

	}

}
