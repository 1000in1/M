package vda5050

import (
	"fmt"

	"github.com/1000in1/m/vda5050/instantactions"
	"github.com/1000in1/m/vda5050/order"
	"github.com/1000in1/m/vda5050/state"
)

type Logger interface {
	INFO(message string)
	ERROR(message string)
}

type VDA5050 struct {
	State    state.State
	TaskId   string
	DeviceId string
	logger   Logger
}

func NewVDA5050(deviceID string) *VDA5050 {
	return &VDA5050{
		State: state.State{

			Errors:     []state.Error{},
			EdgeStates: []state.EdgeState{},
			NodeStates: []*state.NodeState{},
			Loads:      []state.Load{},
		},
		TaskId:   "",
		DeviceId: deviceID,
		logger:   nil,
	}
}

func (v *VDA5050) INFO(message string) {
	if v.logger != nil {
		v.logger.INFO(message)
	}
}

func (v *VDA5050) ERROR(message string) {
	if v.logger != nil {
		v.logger.ERROR(message)
	}
}

func (v *VDA5050) SetLogger(logger Logger) {
	v.logger = logger
}

func (v *VDA5050) UpdateInstantActions(insAction *instantactions.InstantActions) {

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

	json, err := v.State.Marshal()

	return string(json), err

}
