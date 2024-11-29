// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    order, err := UnmarshalOrder(bytes)
//    bytes, err = order.Marshal()

package order

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

func UnmarshalOrder(data []byte) (Order, error) {
	var r Order
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Order) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// The message schema to communicate orders from master control to the AGV.
type Order struct {
	// Directional connection between two nodes. Array of edge objects to be traversed for
	// fulfilling the order. One node is enough for a valid order. Leave edge list empty for
	// that case.
	Edges []Edge `json:"edges,omitempty"`
	// headerId of the message. The headerId is defined per topic and incremented by 1 with each
	// sent (but not necessarily received) message.
	HeaderID *int64 `json:"headerId,omitempty"`
	// Manufacturer of the AGV
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Array of nodes objects to be traversed for fulfilling the order. One node is enough for a
	// valid order. Leave edge list empty for that case.
	Nodes []Node `json:"nodes,omitempty"`
	// Order Identification. This is to be used to identify multiple order messages that belong
	// to the same order.
	OrderID *string `json:"orderId,omitempty"`
	// orderUpdate identification. Is unique per orderId. If an order update is rejected, this
	// field is to be passed in the rejection message.
	OrderUpdateID *int64 `json:"orderUpdateId,omitempty"`
	// Serial number of the AGV.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Timestamp in ISO8601 format (YYYY-MM-DDTHH:mm:ss.ssZ).
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Version of the protocol [Major].[Minor].[Patch]
	Version *string `json:"version,omitempty"`
	// Unique identifier of the zone set that the AGV has to use for navigation or that was used
	// by MC for planning.
	// Optional: Some MC systems do not use zones. Some AGVs do not understand zones. Do not add
	// to message if no zones are used.
	ZoneSetID *string `json:"zoneSetId,omitempty"`

	TaskId *string `json:"taskId,-"`
}

type Edge struct {
	// Array of action objects with detailed information.
	Actions []Action `json:"actions,omitempty"`
	// Sets direction at junctions for line-guided or wire-guided vehicles, to be defined
	// initially (vehicle-individual).
	Direction *string `json:"direction,omitempty"`
	// Additional information on the edge.
	EdgeDescription *string `json:"edgeDescription,omitempty"`
	// Unique edge identification
	EdgeID *string `json:"edgeId,omitempty"`
	// The nodeId of the end node.
	EndNodeID *string `json:"endNodeId,omitempty"`
	// Distance of the path from startNode to endNode in meters.
	// Optional: This value is used by line-guided AGVs to decrease their speed before reaching
	// a stop position.
	Length *float64 `json:"length,omitempty"`
	// Permitted maximum height of the vehicle, including the load, on edge in meters.
	MaxHeight *float64 `json:"maxHeight,omitempty"`
	// Maximum rotation speed in rad/s.
	// Optional: No limit, if not set.
	MaxRotationSpeed *float64 `json:"maxRotationSpeed,omitempty"`
	// Permitted maximum speed on the edge in m/s. Speed is defined by the fastest measurement
	// of the vehicle.
	MaxSpeed *float64 `json:"maxSpeed,omitempty"`
	// Permitted minimal height of the load handling device on the edge in meters
	MinHeight *float64 `json:"minHeight,omitempty"`
	// Orientation of the AGV on the edge. The value orientationType defines if it has to be
	// interpreted relative to the global project specific map coordinate system or tangential
	// to the edge. In case of interpreted tangential to the edge 0.0 = forwards and PI =
	// backwards. Example: orientation Pi/2 rad will lead to a rotation of 90 degrees.
	// If AGV starts in different orientation, rotate the vehicle on the edge to the desired
	// orientation if rotationAllowed is set to True. If rotationAllowed is False, rotate before
	// entering the edge. If that is not possible, reject the order.
	// If no trajectory is defined, apply the rotation to the direct path between the two
	// connecting nodes of the edge. If a trajectory is defined for the edge, apply the
	// orientation to the trajectory.
	Orientation *float64 `json:"orientation,omitempty"`
	// Enum {GLOBALGLOBAL, TANGENTIALTANGENTIAL}:
	// "GLOBAL"- relative to the global project specific map coordinate system;
	// "TANGENTIAL"- tangential to the edge.
	// If not defined, the default value is "TANGENTIAL".
	OrientationType *string `json:"orientationType,omitempty"`
	// True indicates that the edge is part of the base. False indicates that the edge is part
	// of the horizon.
	Released *bool `json:"released,omitempty"`
	// True: rotation is allowed on the edge. False: rotation is not allowed on the edge.
	// Optional: No limit, if not set.
	RotationAllowed *bool `json:"rotationAllowed,omitempty"`
	// Number to track the sequence of nodes and edges in an order and to simplify order
	// updates. The variable sequenceId runs across all nodes and edges of the same order and is
	// reset when a new orderId is issued.
	SequenceID *int64 `json:"sequenceId,omitempty"`
	// The nodeId of the start node.
	StartNodeID *string `json:"startNodeId,omitempty"`
	// Trajectory JSON-object for this edge as a NURBS. Defines the curve, on which the AGV
	// should move between startNode and endNode.
	// Optional: Can be omitted, if AGV cannot process trajectories or if AGV plans its own
	// trajectory.
	Trajectory *Trajectory `json:"trajectory,omitempty"`
}

// Describes an action that the AGV can perform.
type Action struct {
	// Additional information on the action.
	ActionDescription *string `json:"actionDescription,omitempty"`
	// Unique ID to identify the action and map them to the actionState in the state.
	// Suggestion: Use UUIDs.
	ActionID *string `json:"actionId,omitempty"`
	// Array of actionParameter-objects for the indicated action e. g. deviceId, loadId,
	// external Triggers.
	ActionParameters []ActionParameter `json:"actionParameters,omitempty"`
	// Name of action as described in the first column of "Actions and Parameters". Identifies
	// the function of the action.
	ActionType *string `json:"actionType,omitempty"`
	// Regulates if the action is allowed to be executed during movement and/or parallel to
	// other actions.
	// none: action can happen in parallel with others, including movement.
	// soft: action can happen simultaneously with others, but not while moving.
	// hard: no other actions can be performed while this action is running.
	BlockingType *BlockingType `json:"blockingType,omitempty"`
}

type ActionParameter struct {
	// The key of the action parameter.
	Key *string `json:"key,omitempty"`
	// The value of the action parameter
	Value *Value `json:"value"`
}

// Trajectory JSON-object for this edge as a NURBS. Defines the curve, on which the AGV
// should move between startNode and endNode.
// Optional: Can be omitted, if AGV cannot process trajectories or if AGV plans its own
// trajectory.
type Trajectory struct {
	// List of JSON controlPoint objects defining the control points of the NURBS, which
	// includes the beginning and end point.
	ControlPoints []ControlPoint `json:"controlPoints,omitempty"`
	// Defines the number of control points that influence any given point on the curve.
	// Increasing the degree increases continuity. If not defined, the default value is 1.
	Degree *int64 `json:"degree,omitempty"`
	// Sequence of parameter values that determines where and how the control points affect the
	// NURBS curve. knotVector has size of number of control points + degree + 1.
	KnotVector []float64 `json:"knotVector,omitempty"`
}

type ControlPoint struct {
	// The weight, with which this control point pulls on the curve. When not defined, the
	// default will be 1.0.
	Weight *float64 `json:"weight,omitempty"`
	// X coordinate described in the world coordinate system.
	X *float64 `json:"x,omitempty"`
	// Y coordinate described in the world coordinate system.
	Y *float64 `json:"y,omitempty"`
}

type Node struct {
	// Array of actions to be executed on a node. Empty array, if no actions required.
	Actions []Action `json:"actions,omitempty"`
	// Additional information on the node.
	NodeDescription *string `json:"nodeDescription,omitempty"`
	// Unique node identification
	NodeID *string `json:"nodeId,omitempty"`
	// Defines the position on a map in world coordinates. Each floor has its own map. All maps
	// must use the same project specific global origin.
	// Optional for vehicle-types that do not require the node position (e.g., line-guided
	// vehicles).
	NodePosition *NodePosition `json:"nodePosition,omitempty"`
	// True indicates that the node is part of the base. False indicates that the node is part
	// of the horizon.
	Released *bool `json:"released,omitempty"`
	// Number to track the sequence of nodes and edges in an order and to simplify order
	// updates.
	// The main purpose is to distinguish between a node which is passed more than once within
	// one orderId. The variable sequenceId runs across all nodes and edges of the same order
	// and is reset when a new orderId is issued.
	SequenceID *int64 `json:"sequenceId,omitempty"`
}

// Defines the position on a map in world coordinates. Each floor has its own map. All maps
// must use the same project specific global origin.
// Optional for vehicle-types that do not require the node position (e.g., line-guided
// vehicles).
type NodePosition struct {
	// Indicates how big the deviation of theta angle can be.
	// The lowest acceptable angle is theta - allowedDeviationTheta and the highest acceptable
	// angle is theta + allowedDeviationTheta.
	AllowedDeviationTheta *float64 `json:"allowedDeviationTheta,omitempty"`
	// Indicates how exact an AGV has to drive over a node in order for it to count as
	// traversed.
	// If = 0: no deviation is allowed (no deviation means within the normal tolerance of the
	// AGV manufacturer).
	// If > 0: allowed deviation-radius in meters. If the AGV passes a node within the
	// deviation-radius, the node is considered to have been traversed.
	AllowedDeviationXy *float64 `json:"allowedDeviationXy,omitempty"`
	// Additional information on the map.
	MapDescription *string `json:"mapDescription,omitempty"`
	// Unique identification of the map in which the position is referenced.
	// Each map has the same origin of coordinates. When an AGV uses an elevator, e.g., leading
	// from a departure floor to a target floor, it will disappear off the map of the departure
	// floor and spawn in the related lift node on the map of the target floor.
	MapID *string `json:"mapId,omitempty"`
	// Absolute orientation of the AGV on the node.
	// Optional: vehicle can plan the path by itself.
	// If defined, the AGV has to assume the theta angle on this node. If previous edge
	// disallows rotation, the AGV must rotate on the node. If following edge has a differing
	// orientation defined but disallows rotation, the AGV is to rotate on the node to the edges
	// desired rotation before entering the edge.
	Theta *float64 `json:"theta,omitempty"`
	// X-position on the map in reference to the map coordinate system. Precision is up to the
	// specific implementation.
	X *float64 `json:"x,omitempty"`
	// Y-position on the map in reference to the map coordinate system. Precision is up to the
	// specific implementation.
	Y *float64 `json:"y,omitempty"`
}

// Regulates if the action is allowed to be executed during movement and/or parallel to
// other actions.
// none: action can happen in parallel with others, including movement.
// soft: action can happen simultaneously with others, but not while moving.
// hard: no other actions can be performed while this action is running.
type BlockingType string

const (
	Hard BlockingType = "HARD"
	None BlockingType = "NONE"
	Soft BlockingType = "SOFT"
)

// The value of the action parameter
type Value struct {
	AnythingArray []interface{}
	Bool          *bool
	Double        *float64
	String        *string
}

func (x *Value) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Value) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
