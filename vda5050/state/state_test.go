package state

import (
	"testing"
)

func Test_state_Marshal(t *testing.T) {
	state := &State{}

	state.HeaderID = 3233
	state.OrderID = "123-123123"

	state.BatteryState.Charging = false
	state.BatteryState.BatteryHealth = 1
	state.BatteryState.BatteryCharge = 0.5
	state.BatteryState.BatteryVoltage = 12.3
	state.BatteryState.Reach = 9999999

	state.Driving = true

	data, err := state.Marshal()
	if err != nil {
		t.Error(err)
	}

	t.Log(string(data))
}
