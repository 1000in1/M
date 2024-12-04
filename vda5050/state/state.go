// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    state, err := UnmarshalState(bytes)
//    bytes, err = state.Marshal()

package state

import (
	"encoding/json"
	"time"

	"github.com/1000in1/m/vda5050/order"
)

func UnmarshalState(data []byte) (State, error) {
	var r State
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r State) Marshal() ([]byte, error) {

	r.Timestamp = time.Now().Format("2006-01-02T15:04:05.00Z")
	return json.Marshal(r)
}

// all encompassing state of the AGV.
type State struct {
	// Contains a list of the current actions and the actions which are yet to be finished. This
	// may include actions from previous nodes that are still in progress
	// When an action is completed, an updated state message is published with actionStatus set
	// to finished and if applicable with the corresponding resultDescription. The actionStates
	// are kept until a new order is received.
	ActionStates []*ActionState `json:"actionStates"`
	// Defines the position on a map in world coordinates. Each floor has its own map.
	AgvPosition AgvPosition `json:"agvPosition"`
	// Contains all battery-related information.
	BatteryState BatteryState `json:"batteryState"`
	// Used by line guided vehicles to indicate the distance it has been driving past the
	// "lastNodeId".
	// Distance is in meters.
	DistanceSinceLastNode float64 `json:"distanceSinceLastNode,omitempty"`
	// True: indicates that the AGV is driving and/or rotating. Other movements of the AGV
	// (e.g., lift movements) are not included here.
	// False: indicates that the AGV is neither driving nor rotating
	Driving bool `json:"driving"`
	// Array of edgeState-Objects, that need to be traversed for fulfilling the order, empty
	// list if idle.
	EdgeStates []EdgeState `json:"edgeStates"`
	// Array of error-objects. All active errors of the AGV should be in the list. An empty
	// array indicates that the AGV has no active errors.
	Errors []Error `json:"errors"`
	// headerId of the message. The headerId is defined per topic and incremented by 1 with each
	// sent (but not necessarily received) message.
	HeaderID int64 `json:"headerId,omitempty"`
	// Array of info-objects. An empty array indicates, that the AGV has no information. This
	// should only be used for visualization or debugging – it must not be used for logic in
	// master control.
	Information []Information `json:"information,omitempty"`
	// nodeID of last reached node or, if AGV is currently on a node, current node (e.g.,
	// "node7"). Empty string ("") if no lastNodeId is available.
	LastNodeID string `json:"lastNodeId"`
	// sequenceId of the last reached node or, if the AGV is currently on a node, sequenceId of
	// current node.
	// â€œ0â€ if no lastNodeSequenceId is available.
	LastNodeSequenceID int64 `json:"lastNodeSequenceId"`
	// Loads, that are currently handled by the AGV. Optional: If AGV cannot determine load
	// state, leave the array out of the state. If the AGV can determine the load state, but the
	// array is empty, the AGV is considered unloaded.
	Loads []Load `json:"loads"`
	// Manufacturer of the AGV
	Manufacturer string `json:"manufacturer"`
	// True: AGV is almost at the end of the base and will reduce speed if no new base is
	// transmitted. Trigger for master control to send new base
	// False: no base update required.
	NewBaseRequest bool `json:"newBaseRequest"`
	// Array of nodeState-Objects, that need to be traversed for fulfilling the order. Empty
	// list if idle.
	NodeStates []*NodeState `json:"nodeStates"`
	// Current operating mode of the AGV.
	OperatingMode OperatingMode `json:"operatingMode"`
	// Unique order identification of the current order or the previous finished order. The
	// orderId is kept until a new order is received. Empty string ("") if no previous orderId
	// is available.
	OrderID string `json:"orderId"`
	// Order Update Identification to identify that an order update has been accepted by the
	// AGV. "0" if no previous orderUpdateId is available.
	OrderUpdateID int64 `json:"orderUpdateId"`
	// True: AGV is currently in a paused state, either because of the push of a physical button
	// on the AGV or because of an instantAction. The AGV can resume the order.
	// False: The AGV is currently not in a paused state.
	Paused bool `json:"paused"`
	// Contains all safety-related information.
	SafetyState SafetyState `json:"safetyState"`
	// Serial number of the AGV.
	SerialNumber string `json:"serialNumber"`
	// Timestamp in ISO8601 format (YYYY-MM-DDTHH:mm:ss.ssZ).
	Timestamp string `json:"timestamp"`
	// The AGVs velocity in vehicle coordinates
	Velocity Velocity `json:"velocity"`
	// Version of the protocol [Major].[Minor].[Patch]
	Version string `json:"version,omitempty"`
	// Unique ID of the zone set that the AGV currently uses for path planning. Must be the same
	// as the one used in the order, otherwise the AGV is to reject the order.
	// Optional: If the AGV does not use zones, this field can be omitted.
	ZoneSetID string `json:"zoneSetId,omitempty"`
}

type ActionState struct {
	// Additional information on the current action.
	ActionDescription string `json:"actionDescription,omitempty"`
	// Unique actionId
	ActionID string `json:"actionId,omitempty"`
	// WAITING: waiting for the trigger (passing the mode, entering the edge) PAUSED: paused by
	// instantAction or external trigger FAILED: action could not be performed.
	ActionStatus ActionStatus `json:"actionStatus,omitempty"`
	// actionType of the action.
	// Optional: Only for informational or visualization purposes. Order knows the type.
	ActionType string `json:"actionType,omitempty"`
	// Description of the result, e.g., the result of a RFID-read. Errors will be transmitted in
	// errors.
	ResultDescription string `json:"resultDescription,omitempty"`

	//扩展信息
	ActionParameters []order.ActionParameter `json:"-"`

	// Regulates if the action is allowed to be executed during movement and/or parallel to
	// other actions.
	// none: action can happen in parallel with others, including movement.
	// soft: action can happen simultaneously with others, but not while moving.
	// hard: no other actions can be performed while this action is running.
	BlockingType string `json:"-"`

	NodeID string `json:"-"`

	IsSended bool `json:"-"`
}

// Defines the position on a map in world coordinates. Each floor has its own map.
type AgvPosition struct {
	// Value for position deviation range in meters. Optional for vehicles that cannot estimate
	// their deviation, e.g., grid-based localization. Only for logging and visualization
	// purposes.
	DeviationRange float64 `json:"deviationRange,omitempty"`
	// Describes the quality of the localization and therefore, can be used, e.g., by SLAM-AGV
	// to describe how accurate the current position information is.
	// 0.0: position unknown
	// 1.0: position known
	// Optional for vehicles that cannot estimate their localization score.
	// Only for logging and visualization purposes
	LocalizationScore float64 `json:"localizationScore"`
	MapDescription    string  `json:"mapDescription,omitempty"`
	MapID             string  `json:"mapId"`
	// True: position is initialized. False: position is not initizalized.
	PositionInitialized bool    `json:"positionInitialized"`
	Theta               float64 `json:"theta"`
	X                   float64 `json:"x"`
	Y                   float64 `json:"y"`
}

// Contains all battery-related information.
type BatteryState struct {
	// State of Charge in %:
	// If AGV only provides values for good or bad battery levels, these will be indicated as
	// 20% (bad) and 80% (good).
	BatteryCharge float64 `json:"batteryCharge,omitempty"`
	// State of health in percent.
	BatteryHealth float64 `json:"batteryHealth,omitempty"`
	// Battery voltage
	BatteryVoltage float64 `json:"batteryVoltage,omitempty"`
	// True: charging in progress. False: AGV is currently not charging.
	Charging bool `json:"charging"`
	// Estimated reach with current State of Charge in meter.
	Reach float64 `json:"reach,omitempty"`
}

type EdgeState struct {
	// Additional information on the edge.
	EdgeDescription string `json:"edgeDescription,omitempty"`
	// Unique edge identification
	EdgeID string `json:"edgeId,omitempty"`
	// True indicates that the edge is part of the base. False indicates that the edge is part
	// of the horizon.
	Released bool `json:"released,omitempty"`
	// sequenceId of the edge.
	SequenceID int64 `json:"sequenceId,omitempty"`
	// The trajectory is to be communicated as a NURBS and is defined in chapter 6.7
	// Implementation of the Order message.
	// Trajectory segments reach from the point, where the AGV starts to enter the edge to the
	// point where it reports that the next node was traversed.
	Trajectory Trajectory `json:"trajectory,omitempty"`
}

// The trajectory is to be communicated as a NURBS and is defined in chapter 6.7
// Implementation of the Order message.
// Trajectory segments reach from the point, where the AGV starts to enter the edge to the
// point where it reports that the next node was traversed.
type Trajectory struct {
	// List of JSON controlPoint objects defining the control points of the NURBS, which
	// includes the beginning and end point.
	ControlPoints []ControlPoint `json:"controlPoints,omitempty"`
	// Defines the number of control points that influence any given point on the curve.
	// Increasing the degree increases continuity. If not defined, the default value is 1.
	Degree int64 `json:"degree,omitempty"`
	// Sequence of parameter values that determine where and how the control points affect the
	// NURBS curve. knotVector has size of number of control points + degree + 1
	KnotVector []float64 `json:"knotVector,omitempty"`
}

type ControlPoint struct {
	// The weight, with which this control point pulls on the curve.
	// When not defined, the default will be 1.0.
	Weight float64 `json:"weight,omitempty"`
	X      float64 `json:"x,omitempty"`
	Y      float64 `json:"y,omitempty"`
}

type Error struct {
	// Error description.
	ErrorDescription string `json:"errorDescription,omitempty"`
	// WARNING: AGV is ready to start (e.g., maintenance cycle expiration warning). FATAL: AGV
	// is not in running condition, user intervention required (e.g., laser scanner is
	// contaminated).
	ErrorLevel      ErrorLevel       `json:"errorLevel,omitempty"`
	ErrorReferences []ErrorReference `json:"errorReferences,omitempty"`
	// Type/name of error.
	ErrorType string `json:"errorType,omitempty"`
}

// Array of references to identify the source of the error (e.g., headerId, orderId,
// actionId, etc.).
type ErrorReference struct {
	// References the type of reference (e.g., headerId, orderId, actionId, etc.).
	ReferenceKey string `json:"referenceKey,omitempty"`
	// References the value, which belongs to the reference key.
	ReferenceValue string `json:"referenceValue,omitempty"`
}

type Information struct {
	// Info of description.
	InfoDescription string `json:"infoDescription,omitempty"`
	// DEBUG: used for debugging. INFO: used for visualization.
	InfoLevel      InfoLevel       `json:"infoLevel,omitempty"`
	InfoReferences []InfoReference `json:"infoReferences,omitempty"`
	// Type/name of information.
	InfoType string `json:"infoType,omitempty"`
}

// Array of references.
type InfoReference struct {
	// References the type of reference (e.g., headerId, orderId, actionId, etc.).
	ReferenceKey string `json:"referenceKey,omitempty"`
	// References the value, which belongs to the reference key.
	ReferenceValue string `json:"referenceValue,omitempty"`
}

// Load object that describes the load if the AGV has information about it.
type Load struct {
	// Point of reference for the location of the bounding box. The point of reference is always
	// the center of the bounding box bottom surface (at height = 0) and is described in
	// coordinates of the AGV coordinate system.
	BoundingBoxReference BoundingBoxReference `json:"boundingBoxReference,omitempty"`
	// Dimensions of the loads bounding box in meters.
	LoadDimensions LoadDimensions `json:"loadDimensions,omitempty"`
	// Unique identification number of the load (e.g., barcode or RFID). Empty field, if the AGV
	// can identify the load, but did not identify the load yet. Optional, if the AGV cannot
	// identify the load.
	LoadID string `json:"loadId,omitempty"`
	// Indicates, which load handling/carrying unit of the AGV is used, e.g., in case the AGV
	// has multiple spots/positions to carry loads. Optional for vehicles with only one
	// loadPosition.
	LoadPosition string `json:"loadPosition,omitempty"`
	// Type of load.
	LoadType string `json:"loadType,omitempty"`
	// Absolute weight of the load measured in kg.
	Weight float64 `json:"weight,omitempty"`
}

// Point of reference for the location of the bounding box. The point of reference is always
// the center of the bounding box bottom surface (at height = 0) and is described in
// coordinates of the AGV coordinate system.
type BoundingBoxReference struct {
	// Orientation of the loads bounding box. Important for tugger, trains, etc.
	Theta float64 `json:"theta,omitempty"`
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
	Z     float64 `json:"z,omitempty"`
}

// Dimensions of the loads bounding box in meters.
type LoadDimensions struct {
	// Absolute height of the loads bounding box in meter.
	// Optional:
	// Set value only if known.
	Height float64 `json:"height,omitempty"`
	// Absolute length of the loads bounding box in meter.
	Length float64 `json:"length,omitempty"`
	// Absolute width of the loads bounding box in meter.
	Width float64 `json:"width,omitempty"`
}

type NodeState struct {
	// Additional information on the node.
	NodeDescription string `json:"nodeDescription,omitempty"`
	// Unique node identification
	NodeID string `json:"nodeId,omitempty"`
	// Node position. The object is defined in chapter 5.4 Topic: Order (from master control to
	// AGV).
	// Optional:Master control has this information. Can be sent additionally, e.g., for
	// debugging purposes.
	NodePosition NodePosition `json:"nodePosition,omitempty"`
	// True: indicates that the node is part of the base. False: indicates that the node is part
	// of the horizon.
	Released bool `json:"released,omitempty"`
	// sequenceId to discern multiple nodes with same nodeId.
	SequenceID int64 `json:"sequenceId,omitempty"`
}

// Node position. The object is defined in chapter 5.4 Topic: Order (from master control to
// AGV).
// Optional:Master control has this information. Can be sent additionally, e.g., for
// debugging purposes.
type NodePosition struct {
	MapID string  `json:"mapId,omitempty"`
	Theta float64 `json:"theta,omitempty"`
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
}

// Contains all safety-related information.
type SafetyState struct {
	// Acknowledge-Type of eStop: AUTOACK: auto-acknowledgeable e-stop is activated, e.g., by
	// bumper or protective field. MANUAL: e-stop hast to be acknowledged manually at the
	// vehicle. REMOTE: facility e-stop has to be acknowledged remotely. NONE: no e-stop
	// activated.
	EStop EStop `json:"eStop,omitempty"`
	// Protective field violation. True: field is violated. False: field is not violated.
	FieldViolation bool `json:"fieldViolation,omitempty"`
}

// The AGVs velocity in vehicle coordinates
type Velocity struct {
	// The AVGs turning speed around its z axis.
	Omega float64 `json:"omega,omitempty"`
	// The AVGs velocity in its x direction
	Vx float64 `json:"vx,omitempty"`
	// The AVGs velocity in its y direction
	Vy float64 `json:"vy,omitempty"`
}

// WAITING: waiting for the trigger (passing the mode, entering the edge) PAUSED: paused by
// instantAction or external trigger FAILED: action could not be performed.
type ActionStatus string

const (
	Failed       ActionStatus = "FAILED"
	Finished     ActionStatus = "FINISHED"
	Initializing ActionStatus = "INITIALIZING"
	Running      ActionStatus = "RUNNING"
	Waiting      ActionStatus = "WAITING"
)

// WARNING: AGV is ready to start (e.g., maintenance cycle expiration warning). FATAL: AGV
// is not in running condition, user intervention required (e.g., laser scanner is
// contaminated).
type ErrorLevel string

const (
	Fatal   ErrorLevel = "FATAL"
	Warning ErrorLevel = "WARNING"
)

// DEBUG: used for debugging. INFO: used for visualization.
type InfoLevel string

const (
	Debug InfoLevel = "DEBUG"
	Info  InfoLevel = "INFO"
)

// Current operating mode of the AGV.
type OperatingMode string

const (
	Automatic           OperatingMode = "AUTOMATIC"
	OperatingModeMANUAL OperatingMode = "MANUAL"
	Semiautomatic       OperatingMode = "SEMIAUTOMATIC"
	Service             OperatingMode = "SERVICE"
	Teachin             OperatingMode = "TEACHIN"
)

// Acknowledge-Type of eStop: AUTOACK: auto-acknowledgeable e-stop is activated, e.g., by
// bumper or protective field. MANUAL: e-stop hast to be acknowledged manually at the
// vehicle. REMOTE: facility e-stop has to be acknowledged remotely. NONE: no e-stop
// activated.
type EStop string

const (
	Autoack     EStop = "AUTOACK"
	EStopMANUAL EStop = "MANUAL"
	None        EStop = "NONE"
	Remote      EStop = "REMOTE"
)
