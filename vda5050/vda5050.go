package vda5050

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/1000in1/m/logger"
	"github.com/1000in1/m/vda5050/connection"
	"github.com/1000in1/m/vda5050/instantactions"
	"github.com/1000in1/m/vda5050/order"
	"github.com/1000in1/m/vda5050/state"
)

type VDA5050 struct {
	tag      string
	State    state.State
	TaskId   string
	DeviceId string
	logger   logger.LoggerIF
	headerID int64
	lock     sync.Mutex
}

func NewVDA5050(deviceID string, manufacturer string, ver string) *VDA5050 {

	return &VDA5050{
		State: state.State{
			ActionStates: []*state.ActionState{},
			Errors:       []state.Error{},
			EdgeStates:   []state.EdgeState{},
			NodeStates:   []*state.NodeState{},
			Loads:        []state.Load{},
			Version:      ver,
			Manufacturer: manufacturer,
			SerialNumber: deviceID,
		},
		tag:      "VDA5050",
		TaskId:   "",
		DeviceId: deviceID,
		logger:   nil,
		headerID: 0,
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

	v.TaskId = *vda5050_order.TaskId
	v.State.OrderID = *vda5050_order.OrderID
	v.State.OrderUpdateID = *vda5050_order.OrderUpdateID

	for _, node := range vda5050_order.Nodes {

		nodeState := &state.NodeState{
			NodeID:     *node.NodeID,
			SequenceID: *node.SequenceID,
			Released:   *node.Released,
		}

		nodeState.NodePosition.X = *node.NodePosition.X
		nodeState.NodePosition.Y = *node.NodePosition.Y
		nodeState.NodePosition.MapID = *node.NodePosition.MapID

		v.State.NodeStates = append(v.State.NodeStates, nodeState)

		for _, action := range node.Actions {

			actionState := state.ActionState{
				ActionID:         *action.ActionID,
				ActionType:       *action.ActionType,
				BlockingType:     string(*action.BlockingType),
				ActionStatus:     state.Waiting,
				NodeID:           *node.NodeID,
				ActionParameters: action.ActionParameters,
			}

			v.State.ActionStates = append(v.State.ActionStates, &actionState)

		}

	}

	v.INFO(fmt.Sprintf("NodeCount:%d", len(v.State.NodeStates)))

}

func (v *VDA5050) GetNodes() []*state.NodeState {
	return v.State.NodeStates
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
func (v *VDA5050) GetAgvPosition() *state.AgvPosition   { return &v.State.AgvPosition }

func (v *VDA5050) GetSafetyState() *state.SafetyState { return &v.State.SafetyState }

func (v *VDA5050) ChangeActionStatus(actionId string, status state.ActionStatus) bool {
	v.lock.Lock()
	defer v.lock.Unlock()

	for _, action := range v.State.ActionStates {
		if action.ActionID == actionId {
			//last := action.ActionStatus
			action.ActionStatus = status
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
