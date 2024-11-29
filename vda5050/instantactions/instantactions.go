// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    instantActions, err := UnmarshalInstantActions(bytes)
//    bytes, err = instantActions.Marshal()

package instantactions

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"github.com/1000in1/m/vda5050/order"
)

func UnmarshalInstantActions(data []byte) (InstantActions, error) {
	var r InstantActions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstantActions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// JSON Schema for publishing instantActions that the AGV is to execute as soon as they
// arrive.
type InstantActions struct {
	Actions []Action `json:"actions,omitempty"`
	// headerId of the message. The headerId is defined per topic and incremented by 1 with each
	// sent (but not necessarily received) message.
	HeaderID *int64 `json:"headerId,omitempty"`
	// Manufacturer of the AGV
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Serial number of the AGV.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Timestamp in ISO8601 format (YYYY-MM-DDTHH:mm:ss.ssZ).
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Version of the protocol [Major].[Minor].[Patch]
	Version *string `json:"version,omitempty"`
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
	ActionParameters []order.ActionParameter `json:"actionParameters,omitempty"`
	// Name of action as described in the first column of "Actions and Parameters". Identifies
	// the function of the action.
	ActionType *string `json:"actionType,omitempty"`
	// Regulates if the action is allowed to be executed during movement and/or parallel to
	// other actions.
	// none: action can happen in parallel with others, including movement.
	// soft: action can happen simultaneously with others, but not while moving.
	// hard: no other actions can be performed while this action is running.
	BlockingType *order.BlockingType `json:"blockingType,omitempty"`
}

type ActionParameter struct {
	// The key of the action parameter.
	Key *string `json:"key,omitempty"`
	// The value of the action parameter
	Value *Value `json:"value"`
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
