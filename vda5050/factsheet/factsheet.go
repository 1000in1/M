// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    factsheet, err := UnmarshalFactsheet(bytes)
//    bytes, err = factsheet.Marshal()

package factsheet

import (
	"encoding/json"
	"time"
)

func UnmarshalFactsheet(data []byte) (Factsheet, error) {
	var r Factsheet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Factsheet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// The factsheet provides basic information about a specific AGV type series. This
// information allows comparison of different AGV types and can be applied for the planning,
// dimensioning and simulation of an AGV system. The factsheet also includes information
// about AGV communication interfaces which are required for the integration of an AGV type
// series into a VD[M]A-5050-compliant master control.
type Factsheet struct {
	// Detailed definition of AGV geometry
	AgvGeometry *AgvGeometry `json:"agvGeometry,omitempty"`
	// header ID of the message. The headerId is defined per topic and incremented by 1 with
	// each sent (but not necessarily received) message.
	HeaderID *int64 `json:"headerId,omitempty"`
	// Abstract specification of load capabilities
	LoadSpecification *LoadSpecification `json:"loadSpecification,omitempty"`
	// Detailed specification of localization
	LocalizationParameters *int64 `json:"localizationParameters,omitempty"`
	// Manufacturer of the AGV
	Manufacturer *string `json:"manufacturer,omitempty"`
	// These parameters specify the basic physical properties of the AGV
	PhysicalParameters *PhysicalParameters `json:"physicalParameters,omitempty"`
	// Supported features of VDA5050 protocol
	ProtocolFeatures *ProtocolFeatures `json:"protocolFeatures,omitempty"`
	// This JSON-object describes the protocol limitations of the AGV. If a parameter is not
	// defined or set to zero then there is no explicit limit for this parameter.
	ProtocolLimits *ProtocolLimits `json:"protocolLimits,omitempty"`
	// Serial number of the AGV
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Timestamp in ISO8601 format (YYYY-MM-DDTHH:mm:ss.ssZ).
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// These parameters generally specify the class and the capabilities of the AGV
	TypeSpecification *TypeSpecification `json:"typeSpecification,omitempty"`
	// Version of the VD[M]A-5050 protocol [Major].[Minor].[Patch] (e.g. 1.3.2)
	Version *string `json:"version,omitempty"`
}

// Detailed definition of AGV geometry
type AgvGeometry struct {
	Envelopes2D []Envelopes2D `json:"envelopes2d,omitempty"`
	// list of AGV-envelope curves in 3D (german: „Hüllkurven“)
	Envelopes3D []Envelopes3D `json:"envelopes3d,omitempty"`
	// list of wheels, containing wheel-arrangement and geometry
	WheelDefinitions []WheelDefinition `json:"wheelDefinitions,omitempty"`
}

type Envelopes2D struct {
	// free text: description of envelope curve set
	Description *string `json:"description,omitempty"`
	// envelope curve as a x/y-polygon polygon is assumed as closed and must be
	// non-self-intersecting
	PolygonPoints []PolygonPoint `json:"polygonPoints,omitempty"`
	// name of the envelope curve set
	Set *string `json:"set,omitempty"`
}

type PolygonPoint struct {
	// x-position of polygon-point
	X *float64 `json:"x,omitempty"`
	// y-position of polygon-point
	Y *float64 `json:"y,omitempty"`
}

type Envelopes3D struct {
	// 3D-envelope curve data, format specified in ‚format‘
	Data map[string]interface{} `json:"data,omitempty"`
	// free text: description of envelope curve set
	Description *int64 `json:"description,omitempty"`
	// format of data e.g. DXF
	Format *string `json:"format,omitempty"`
	// name of the envelope curve set
	Set *string `json:"set,omitempty"`
	// protocol and url-definition for downloading the 3D-envelope curve data e.g.
	// ftp://xxx.yyy.com/ac4dgvhoif5tghji
	URL *string `json:"url,omitempty"`
}

type WheelDefinition struct {
	// nominal displacement of the wheel’s center to the rotation point (necessary for caster
	// wheels). If the parameter is not defined, it is assumed to be 0
	CenterDisplacement *float64 `json:"centerDisplacement,omitempty"`
	// free text: can be used by the manufacturer to define constraints
	Constraints *string `json:"constraints,omitempty"`
	// nominal diameter of wheel
	Diameter *float64 `json:"diameter,omitempty"`
	// True: wheel is actively driven (de: angetrieben)
	IsActiveDriven *bool `json:"isActiveDriven,omitempty"`
	// True: wheel is actively steered (de: aktiv gelenkt)
	IsActiveSteered *bool     `json:"isActiveSteered,omitempty"`
	Position        *Position `json:"position,omitempty"`
	// wheel type. DRIVE, CASTER, FIXED, MECANUM
	Type *Type `json:"type,omitempty"`
	// nominal width of wheel
	Width *float64 `json:"width,omitempty"`
}

type Position struct {
	// orientation of wheel in AGV-coordinate system Necessary for fixed wheels
	Theta *float64 `json:"theta,omitempty"`
	// [m] x-position in AGV-coordinate system
	X *float64 `json:"x,omitempty"`
	// y-position in AGV-coordinate system
	Y *float64 `json:"y,omitempty"`
}

// Abstract specification of load capabilities
type LoadSpecification struct {
	// list of load positions / load handling devices. This lists contains the valid values for
	// the oarameter “state.loads[].loadPosition” and for the action parameter “lhd” of the
	// actions pick and drop. If this list doesn’t exist or is empty, the AGV has no load
	// handling device.
	LoadPositions []string `json:"loadPositions,omitempty"`
	// list of load-sets that can be handled by the AGV
	LoadSets []LoadSet `json:"loadSets,omitempty"`
}

type LoadSet struct {
	// maximum allowed acceleration for this load-type and –weight
	AgvAccelerationLimit *float64 `json:"agvAccelerationLimit,omitempty"`
	// maximum allowed deceleration for this load-type and –weight
	AgvDecelerationLimit *float64 `json:"agvDecelerationLimit,omitempty"`
	// maximum allowed speed for this load-type and –weight
	AgvSpeedLimit *float64 `json:"agvSpeedLimit,omitempty"`
	// bounding box reference as defined in parameter loads[] in state-message
	BoundingBoxReference *BoundingBoxReference `json:"boundingBoxReference,omitempty"`
	// free text description of the load handling set
	Description *float64 `json:"description,omitempty"`
	// approx. time for dropping the load
	DropTime       *float64        `json:"dropTime,omitempty"`
	LoadDimensions *LoadDimensions `json:"loadDimensions,omitempty"`
	// list of load positions btw. load handling devices, this load-set is valid for. If this
	// parameter does not exist or is empty, this load-set is valid for all load handling
	// devices on this AGV.
	LoadPositions []string `json:"loadPositions,omitempty"`
	// type of load e.g. EPAL, XLT1200, ….
	LoadType *string `json:"loadType,omitempty"`
	// maximum allowed depth for this load-type and –weight. references to boundingBoxReference
	MaxLoadhandlingDepth *float64 `json:"maxLoadhandlingDepth,omitempty"`
	// maximum allowed height for handling of this load-type and –weight. references to
	// boundingBoxReference
	MaxLoadhandlingHeight *float64 `json:"maxLoadhandlingHeight,omitempty"`
	// maximum allowed tilt for this load-type and –weight
	MaxLoadhandlingTilt *float64 `json:"maxLoadhandlingTilt,omitempty"`
	// maximum weight of loadtype
	MaxWeigth *float64 `json:"maxWeigth,omitempty"`
	// minimum allowed depth for this load-type and –weight. references to boundingBoxReference
	MinLoadhandlingDepth *float64 `json:"minLoadhandlingDepth,omitempty"`
	// minimum allowed height for handling of this load-type and –weight. References to
	// boundingBoxReference
	MinLoadhandlingHeight *float64 `json:"minLoadhandlingHeight,omitempty"`
	// minimum allowed tilt for this load-type and –weight
	MinLoadhandlingTilt *float64 `json:"minLoadhandlingTilt,omitempty"`
	// approx. time for picking up the load
	PickTime *float64 `json:"pickTime,omitempty"`
	// Unique name of the load set, e.g. DEFAULT, SET1, ...
	SetName *string `json:"setName,omitempty"`
}

// bounding box reference as defined in parameter loads[] in state-message
type BoundingBoxReference struct {
	// Orientation of the loads bounding box. Important for tugger trains, etc.
	Theta *int64 `json:"theta,omitempty"`
	// x-coordinate of the point of reference.
	X *float64 `json:"x,omitempty"`
	// y-coordinate of the point of reference.
	Y *float64 `json:"y,omitempty"`
	// z-coordinate of the point of reference.
	Z *float64 `json:"z,omitempty"`
}

type LoadDimensions struct {
	// Absolute height of the load´s bounding box. Optional: Set value only if known.
	Height *float64 `json:"height,omitempty"`
	// Absolute length of the load´s bounding box.
	Length *float64 `json:"length,omitempty"`
	// Absolute width of the load´s bounding bo
	Width *float64 `json:"width,omitempty"`
}

// These parameters specify the basic physical properties of the AGV
type PhysicalParameters struct {
	// maximum acceleration with maximum load
	AccelerationMax *float64 `json:"accelerationMax,omitempty"`
	// maximum deceleration with maximum load
	DecelerationMax *float64 `json:"decelerationMax,omitempty"`
	// maximum height of AGV
	HeightMax *float64 `json:"heightMax,omitempty"`
	// minimum height of AGV
	HeightMin *float64 `json:"heightMin,omitempty"`
	// length of AGV
	Length *float64 `json:"length,omitempty"`
	// maximum speed of the AGV
	SpeedMax *float64 `json:"speedMax,omitempty"`
	// minimal controlled continuous speed of the AGV
	SpeedMin *float64 `json:"speedMin,omitempty"`
	// width of AGV
	Width *float64 `json:"width,omitempty"`
}

// Supported features of VDA5050 protocol
type ProtocolFeatures struct {
	// list of all actions with parameters supported by this AGV. This includes standard actions
	// specified in VDA5050 and manufacturer-specific actions
	AgvActions []AgvAction `json:"agvActions,omitempty"`
	// list of supported and/or required optional parameters. Optional parameters, that are not
	// listed here, are assumed to be not supported by the AGV.
	OptionalParameters []OptionalParameter `json:"optionalParameters,omitempty"`
}

type AgvAction struct {
	// free text: description of the action
	ActionDescription *string `json:"actionDescription,omitempty"`
	// list of parameters. if not defined, the action has no parameters
	ActionParameters []ActionParameter `json:"actionParameters,omitempty"`
	// list of allowed scopes for using this action-type. INSTANT: usable as instantAction,
	// NODE: usable on nodes, EDGE: usable on edges.
	ActionScopes []ActionScope `json:"actionScopes,omitempty"`
	// unique actionType corresponding to action.actionType
	ActionType *string `json:"actionType,omitempty"`
	// free text: description of the resultDescription
	ResultDescription *string `json:"resultDescription,omitempty"`
}

type ActionParameter struct {
	// free text: description of the parameter
	Description *string `json:"description,omitempty"`
	// True: optional parameter
	IsOptional *bool `json:"isOptional,omitempty"`
	// key-String for Parameter
	Key *string `json:"key,omitempty"`
	// data type of Value, possible data types are: BOOL, NUMBER, INTEGER, FLOAT, STRING,
	// OBJECT, ARRAY
	ValueDataType *ValueDataType `json:"valueDataType,omitempty"`
}

type OptionalParameter struct {
	// free text. Description of optional parameter. E.g. Reason, why the optional parameter
	// ‚direction‘ is necessary for this AGV-type and which values it can contain. The parameter
	// ‘nodeMarker’ must contain unsigned interger-numbers only. Nurbs-Support is limited to
	// straight lines and circle segments.
	Description *string `json:"description,omitempty"`
	// full name of optional parameter, e.g. “order.nodes.nodePosition.allowedDeviationTheta”
	Parameter *string `json:"parameter,omitempty"`
	// type of support for the optional parameter, the following values are possible: SUPPORTED:
	// optional parameter is supported like specified. REQUIRED: optional parameter is required
	// for proper AGV-operation.
	Support *Support `json:"support,omitempty"`
}

// This JSON-object describes the protocol limitations of the AGV. If a parameter is not
// defined or set to zero then there is no explicit limit for this parameter.
type ProtocolLimits struct {
	// maximum lengths of arrays
	MaxArrayLens map[string]interface{} `json:"maxArrayLens,omitempty"`
	// maximum lengths of strings
	MaxStringLens *MaxStringLens `json:"maxStringLens,omitempty"`
	// timing information
	Timing *Timing `json:"timing,omitempty"`
}

// maximum lengths of strings
type MaxStringLens struct {
	// maximum length of ENUM- and Key-Strings. Affected parameters: action.actionType,
	// action.blockingType, edge.direction, actionParameter.key, state.operatingMode,
	// load.loadPosition, load.loadType, actionState.actionStatus, error.errorType,
	// error.errorLevel, errorReference.referenceKey, info.infoType, info.infoLevel,
	// safetyState.eStop, connection.connectionState
	EnumLen *int64 `json:"enumLen,omitempty"`
	// maximum length of ID-Strings. Affected parameters: order.orderId, order.zoneSetId,
	// node.nodeId, nodePosition.mapId, action.actionId, edge.edgeId, edge.startNodeId,
	// edge.endNodeId
	IDLen *int64 `json:"idLen,omitempty"`
	// If true ID-strings need to contain numerical values only
	IDNumericalOnly *bool `json:"idNumericalOnly,omitempty"`
	// maximum length of loadId Strings
	LoadIDLen *int64 `json:"loadIdLen,omitempty"`
	// maximum MQTT Message length
	MsgLen *int64 `json:"msgLen,omitempty"`
	// maximum length of all other parts in MQTT-topics. Affected parameters: order.timestamp,
	// order.version, order.manufacturer, instantActions.timestamp, instantActions.version,
	// instantActions.manufacturer, state.timestamp, state.version, state.manufacturer,
	// visualization.timestamp, visualization.version, visualization.manufacturer,
	// connection.timestamp, connection.version, connection.manufacturer
	TopicElemLen *int64 `json:"topicElemLen,omitempty"`
	// maximum length of serial-number part in MQTT-topics. Affected Parameters:
	// order.serialNumber, instantActions.serialNumber, state.SerialNumber,
	// visualization.serialNumber, connection.serialNumber
	TopicSerialLen *int64 `json:"topicSerialLen,omitempty"`
}

// timing information
type Timing struct {
	// default interval for sending state-messages if not defined, the default value from the
	// main document is used
	DefaultStateInterval *float64 `json:"defaultStateInterval,omitempty"`
	// minimum interval sending order messages to the AGV
	MinOrderInterval *float64 `json:"minOrderInterval,omitempty"`
	// minimum interval for sending state-messages
	MinStateInterval *float64 `json:"minStateInterval,omitempty"`
	// default interval for sending messages on visualization topic
	VisualizationInterval *float64 `json:"visualizationInterval,omitempty"`
}

// These parameters generally specify the class and the capabilities of the AGV
type TypeSpecification struct {
	// Simplified description of AGV class.
	AgvClass *AgvClass `json:"agvClass,omitempty"`
	// simplified description of AGV kinematics-type.
	AgvKinematic *AgvKinematic `json:"agvKinematic,omitempty"`
	// simplified description of localization type
	LocalizationTypes []LocalizationType `json:"localizationTypes,omitempty"`
	// maximum loadable mass
	MaxLoadMass *float64 `json:"maxLoadMass,omitempty"`
	// List of path planning types supported by the AGV, sorted by priority
	NavigationTypes []NavigationType `json:"navigationTypes,omitempty"`
	// Free text human readable description of the AGV type series
	SeriesDescription *string `json:"seriesDescription,omitempty"`
	// Free text generalized series name as specified by manufacturer
	SeriesName *string `json:"seriesName,omitempty"`
}

// wheel type. DRIVE, CASTER, FIXED, MECANUM
type Type string

const (
	Caster  Type = "CASTER"
	Drive   Type = "DRIVE"
	Fixed   Type = "FIXED"
	Mecanum Type = "MECANUM"
)

// data type of Value, possible data types are: BOOL, NUMBER, INTEGER, FLOAT, STRING,
// OBJECT, ARRAY
type ValueDataType string

const (
	Array   ValueDataType = "ARRAY"
	Bool    ValueDataType = "BOOL"
	Float   ValueDataType = "FLOAT"
	Integer ValueDataType = "INTEGER"
	Number  ValueDataType = "NUMBER"
	Object  ValueDataType = "OBJECT"
	String  ValueDataType = "STRING"
)

type ActionScope string

const (
	Edge    ActionScope = "EDGE"
	Instant ActionScope = "INSTANT"
	Node    ActionScope = "NODE"
)

// type of support for the optional parameter, the following values are possible: SUPPORTED:
// optional parameter is supported like specified. REQUIRED: optional parameter is required
// for proper AGV-operation.
type Support string

const (
	Required  Support = "REQUIRED"
	Supported Support = "SUPPORTED"
)

// Simplified description of AGV class.
type AgvClass string

const (
	Carrier  AgvClass = "CARRIER"
	Conveyor AgvClass = "CONVEYOR"
	Forklift AgvClass = "FORKLIFT"
	Tugger   AgvClass = "TUGGER"
)

// simplified description of AGV kinematics-type.
type AgvKinematic string

const (
	Diff       AgvKinematic = "DIFF"
	Omni       AgvKinematic = "OMNI"
	Threewheel AgvKinematic = "THREEWHEEL"
)

type LocalizationType string

const (
	Dmc       LocalizationType = "DMC"
	Grid      LocalizationType = "GRID"
	Natural   LocalizationType = "NATURAL"
	RFID      LocalizationType = "RFID"
	Reflector LocalizationType = "REFLECTOR"
	Spot      LocalizationType = "SPOT"
)

type NavigationType string

const (
	Autonomous          NavigationType = "AUTONOMOUS"
	PhysicalLindeGuided NavigationType = "PHYSICAL_LINDE_GUIDED"
	VirtualLineGuided   NavigationType = "VIRTUAL_LINE_GUIDED"
)
