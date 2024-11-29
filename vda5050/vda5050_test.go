package vda5050

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/1000in1/m/vda5050/instantactions"
	"github.com/1000in1/m/vda5050/order"
	"github.com/1000in1/m/vda5050/state"
)

func vda5050_order() (string, error) {
	var json_str = `{  "headerId" : 1903,  "timestamp" : "2024-11-29T05:23:29.044972Z",  "version" : "2.0.0",  "manufacturer" : "autoxing",  "serialNumber" : "L382407a03613Wh",  "orderId" : "Park-01JDV5ZHY0ZK7Y8JE59JZ9F1JN-0",  "orderUpdateId" : 0,  "nodes" : [ {    "nodeId" : "n_22b25ddc4dcd7a2bd822d1e2e7f82470",    "sequenceId" : 0,    "released" : true,    "actions" : [ ],    "nodePosition" : {      "x" : 18.897,      "y" : -29.844,      "mapId" : "673d98a136bbc50b54d4381e",      "allowedDeviationXY" : 0.3638997033058943,      "allowedDeviationTheta" : 3.141592653589793    }  }, {    "nodeId" : "673d98d4f2d42efb5966d862",    "sequenceId" : 2,    "released" : true,    "actions" : [ ],    "nodePosition" : {      "x" : 17.869,      "y" : -28.355,      "mapId" : "673d98a136bbc50b54d4381e",      "theta" : 2.349038576343302,      "allowedDeviationXY" : 0.3,      "allowedDeviationTheta" : 3.141592653589793    }  } ],  "edges" : [ {    "edgeId" : "l_8c1720d9eb300cf151d387799f9a5e55",    "sequenceId" : 1,    "released" : true,    "startNodeId" : "n_22b25ddc4dcd7a2bd822d1e2e7f82470",    "endNodeId" : "673d98d4f2d42efb5966d862",    "actions" : [ ],    "maxSpeed" : 1.0  } ],  "routeIndex" : 0,  "routePoints" : [ "n_22b25ddc4dcd7a2bd822d1e2e7f82470", "673d98d4f2d42efb5966d862" ],  "routePaths" : [ {    "x" : 18897,    "y" : -29844,    "z" : 0  }, {    "x" : 17869,    "y" : -28355,    "z" : 0  } ],  "destinationName" : "673d98d4f2d42efb5966d862",  "allocatedResources" : [ "n_22b25ddc4dcd7a2bd822d1e2e7f82470", "673d98d4f2d42efb5966d862" ],  "finalMovement" : true,  "finalOrder" : true,  "taskId" : "Park-01JDV5ZHY0ZK7Y8JE59JZ9F1JN"}`

	json_str = `{  "headerId" : 2886,  "timestamp" : "2024-11-29T06:35:37.027011Z",  "version" : "2.0.0",  "manufacturer" : "autoxing",  "serialNumber" : "10SIMdev000C-04",  "orderId" : "AX-A-B1-925244bc8038419c-1",  "orderUpdateId" : 14,  "nodes" : [ {    "nodeId" : "n_d3de58669fb134c017f37f6bd07585d9",    "sequenceId" : 28,    "released" : true,    "actions" : [ ],    "nodePosition" : {      "x" : 17.798,      "y" : -39.255,      "mapId" : "6513c9d1af74e21e16a28fc1",      "allowedDeviationXY" : 0.3,      "allowedDeviationTheta" : 3.141592653589793    }  }, {    "nodeId" : "67160164405904345af8f92e",    "sequenceId" : 30,    "released" : true,    "actions" : [ {      "actionType" : "Drop",      "actionId" : "Order_destination_action_14",      "blockingType" : "SOFT",      "actionParameters" : [ {        "key" : "actions",        "value" : "[{\"type\":48,\"data\":{}}]"      } ]    } ],    "nodePosition" : {      "x" : 17.806,      "y" : -37.887,      "mapId" : "6513c9d1af74e21e16a28fc1",      "theta" : 1.5707963267948966,      "allowedDeviationXY" : 0.3,      "allowedDeviationTheta" : 3.141592653589793    }  } ],  "edges" : [ {    "edgeId" : "reverse_l_c682ae492bcf7ae9dbe21fc4807ccf70",    "sequenceId" : 29,    "released" : true,    "startNodeId" : "n_d3de58669fb134c017f37f6bd07585d9",    "endNodeId" : "67160164405904345af8f92e",    "actions" : [ ],    "maxSpeed" : 1.0  } ],  "routeIndex" : 14,  "routePoints" : [ "6715feff60fdac1f872531b6", "n_fe365511d3e062c5c8160c5a7d6b6751", "n_d7637890c812a9ac53ed578904966f0c", "6731b1b66048b8b3d1687701", "n_f025bd640e54591d8cd1ece899f37027", "n_94462ffc573cf2f7a3f1a13e4e25d7c3", "n_51b07a4e3286525aa3109997b5e2dd87", "n_5befbfa5b7418e3c5118d1cc35ffc01f", "n_b91f48fedde659d775c79c1947537fbb", "n_08950f9d595d0225c1c04660ab6ed592", "6731b1dc38a286c6399f26fe", "n_c38e7246f1649ea559ee34fe0959bad9", "n_ea4c070031ba7d9409b469d303ace432", "n_92f7257de39cc5746956f978d001f8a1", "n_d3de58669fb134c017f37f6bd07585d9", "67160164405904345af8f92e" ],  "routePaths" : [ {    "x" : 21395,    "y" : -51375,    "z" : 0  }, {    "x" : 21398,    "y" : -50023,    "z" : 0  }, {    "x" : 21410,    "y" : -48806,    "z" : 0  }, {    "x" : 23156,    "y" : -48815,    "z" : 0  }, {    "x" : 23173,    "y" : -47393,    "z" : 0  }, {    "x" : 23178,    "y" : -46201,    "z" : 0  }, {    "x" : 23183,    "y" : -45109,    "z" : 0  }, {    "x" : 23188,    "y" : -43852,    "z" : 0  }, {    "x" : 23186,    "y" : -42761,    "z" : 0  }, {    "x" : 23198,    "y" : -41633,    "z" : 0  }, {    "x" : 23192,    "y" : -40462,    "z" : 0  }, {    "x" : 21483,    "y" : -40447,    "z" : 0  }, {    "x" : 19678,    "y" : -40439,    "z" : 0  }, {    "x" : 17810,    "y" : -40454,    "z" : 0  }, {    "x" : 17798,    "y" : -39255,    "z" : 0  }, {    "x" : 17806,    "y" : -37887,    "z" : 0  } ],  "destinationName" : "67160164405904345af8f92e",  "allocatedResources" : [ "n_92f7257de39cc5746956f978d001f8a1", "n_d3de58669fb134c017f37f6bd07585d9", "67160164405904345af8f92e" ],  "finalMovement" : true,  "finalOrder" : true,  "taskId" : "AX-A-B1-925244bc8038419c"}`

	vda5050Order, err := order.UnmarshalOrder([]byte(json_str))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vda5050Order)

	vda5050S := NewVDA5050("robot_test_111")

	vda5050S.UpdateOrder(&vda5050Order)

	data, err := vda5050S.GetStateJsonString()

	fmt.Println(vda5050S.TaskId)

	return data, err
}

func Test_vda5050_order(t *testing.T) {

	data, err := vda5050_order()
	if err != nil {
		t.Error(err)
	}
	t.Log(data)

}

// `
// recv: uagv/v2/autoxing/10SIMdev000C-04/instantActions {
//   "headerId" : 861,
//   "timestamp" : "2024-11-29T10:00:23.314731Z",
//   "version" : "2.0.0",
//   "manufacturer" : "autoxing",
//   "serialNumber" : "10SIMdev000C-04",
//   "actions" : [ {
//     "actionType" : "updateLastNodeId",
//     "actionId" : "8d75c857-10ae-4310-9d58-49c2b574f561",
//     "blockingType" : "NONE",
//     "actionParameters" : [ {
//       "key" : "nodeId",
//       "value" : "n_badc902d08c298430785cc3cdf1ddf53"
//     }, {
//       "key" : "isManual",
//       "value" : false
//     } ]
//   } ]
// }

// recv: uagv/v2/autoxing/10SIMdev000C-04/instantActions {
//   "headerId" : 862,
//   "timestamp" : "2024-11-29T10:00:23.332412Z",
//   "version" : "2.0.0",
//   "manufacturer" : "autoxing",
//   "serialNumber" : "10SIMdev000C-04",
//   "actions" : [ {
//     "actionType" : "updateMap",
//     "actionId" : "ec8b67c6-67ff-4d99-8294-57eaf51bada6",
//     "blockingType" : "HARD",
//     "actionDescription" : "update current map",
//     "actionParameters" : [ ]
//   } ]
// }

// recv: uagv/v2/autoxing/10SIMdev000C-04/instantActions {
//   "headerId" : 863,
//   "timestamp" : "2024-11-29T10:00:27.441267Z",
//   "version" : "2.0.0",
//   "manufacturer" : "autoxing",
//   "serialNumber" : "10SIMdev000C-04",
//   "actions" : [ {
//     "actionType" : "resumeTask",
//     "actionId" : "37c3601c-cbd9-4789-bd0f-6b43c68db506",
//     "blockingType" : "HARD",
//     "actionDescription" : "resume the current task",
//     "actionParameters" : [ ]
//   } ]
// }
// recv: uagv/v2/autoxing/10SIMdev000C-04/instantActions {
//   "headerId" : 864,
//   "timestamp" : "2024-11-29T10:00:28.820020Z",
//   "version" : "2.0.0",
//   "manufacturer" : "autoxing",
//   "serialNumber" : "10SIMdev000C-04",
//   "actions" : [ {
//     "actionType" : "startJacking",
//     "actionId" : "2b0d8578-c8b0-4981-90ea-0e842d9aea62",
//     "blockingType" : "HARD",
//     "actionDescription" : "pick the shelves",
//     "actionParameters" : [ {
//       "key" : "type",
//       "value" : 0
//     } ]
//   } ]
// }

// recv: uagv/v2/autoxing/10SIMdev000C-04/instantActions {
//   "headerId" : 865,
//   "timestamp" : "2024-11-29T10:00:29.385138Z",
//   "version" : "2.0.0",
//   "manufacturer" : "autoxing",
//   "serialNumber" : "10SIMdev000C-04",
//   "actions" : [ {
//     "actionType" : "startJacking",
//     "actionId" : "a118c742-9532-44fc-ab7e-338cfb1e781e",
//     "blockingType" : "HARD",
//     "actionDescription" : "drop the shelves",
//     "actionParameters" : [ {
//       "key" : "type",
//       "value" : 1
//     } ]
//   } ]
// }

// recv: uagv/v2/autoxing/10SIMdev000C-04/instantActions {
//   "headerId" : 866,
//   "timestamp" : "2024-11-29T10:00:30.268524Z",
//   "version" : "2.0.0",
//   "manufacturer" : "autoxing",
//   "serialNumber" : "10SIMdev000C-04",
//   "actions" : [ {
//     "actionType" : "updateLastNodeId",
//     "actionId" : "618bef6d-f675-4343-96ae-8a421d549ed5",
//     "blockingType" : "NONE",
//     "actionParameters" : [ {
//       "key" : "nodeId",
//       "value" : "6743dc5fa3858fdda6ffe926"
//     }, {
//       "key" : "isManual",
//       "value" : false
//     } ]
//   } ]
// }

// recv: uagv/v2/autoxing/10SIMdev000C-05/instantActions {
//   "headerId" : 905,
//   "timestamp" : "2024-11-29T10:00:34.172682Z",
//   "version" : "2.0.0",
//   "manufacturer" : "autoxing",
//   "serialNumber" : "10SIMdev000C-05",
//   "actions" : [ {
//     "actionType" : "updateLastNodeId",
//     "actionId" : "2c879c22-66b6-4ff7-aff1-0e774aee4e89",
//     "blockingType" : "NONE",
//     "actionParameters" : [ {
//       "key" : "nodeId",
//       "value" : "6731b182088f262df9845fe2"
//     }, {
//       "key" : "isManual",
//       "value" : false
//     } ]
//   } ]
// }

// `
func testInstantActions(t *testing.T, vda5050s *VDA5050, json_str string) string {

	vda5050InstantActions, err := instantactions.UnmarshalInstantActions([]byte(json_str))

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	vda5050s.UpdateInstantActions(&vda5050InstantActions)

	data, err := vda5050s.State.Marshal()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	return string(data)

}

func vda5050_load_data(t *testing.T) *VDA5050 {

	vda5050s := NewVDA5050("robor_32333")

	var json_str = `{  "headerId" : 365,  "timestamp" : "2024-11-29T07:49:11.964946Z",  "version" : "2.0.0",  "manufacturer" : "autoxing",  "serialNumber" : "10SIMdev000C-05",  "actions" : [ {    "actionType" : "stopPause",    "actionId" : "1b8d94df-963d-4691-9174-434bd0ec1562",    "blockingType" : "NONE"  } ]}`

	json_str = `{  "headerId" : 366,  "timestamp" : "2024-11-29T07:49:13.580231Z",  "version" : "2.0.0",  "manufacturer" : "autoxing",  "serialNumber" : "10SIMdev000C-05",  "actions" : [ {    "actionType" : "startPause",    "actionId" : "c75cf5c5-953e-43e5-aa53-86f8959853d4",    "blockingType" : "NONE"  } ]}`

	_ = testInstantActions(t, vda5050s, json_str)

	json_str = `{  "headerId" : 905,  "timestamp" : "2024-11-29T10:00:34.172682Z",  "version" : "2.0.0",  "manufacturer" : "autoxing",  "serialNumber" : "10SIMdev000C-05",  "actions" : [ {    "actionType" : "updateLastNodeId",    "actionId" : "2c879c22-66b6-4ff7-aff1-0e774aee4e89",    "blockingType" : "NONE",    "actionParameters" : [ {      "key" : "nodeId",      "value" : "6731b182088f262df9845fe2"    }, {      "key" : "isManual",      "value" : false    } ]  } ]}`
	_ = testInstantActions(t, vda5050s, json_str)

	json_str = `{  "headerId" : 865,  "timestamp" : "2024-11-29T10:00:29.385138Z",  "version" : "2.0.0",  "manufacturer" : "autoxing",  "serialNumber" : "10SIMdev000C-04",  "actions" : [ {    "actionType" : "startJacking",    "actionId" : "a118c742-9532-44fc-ab7e-338cfb1e781e",    "blockingType" : "HARD",    "actionDescription" : "drop the shelves",    "actionParameters" : [ {      "key" : "type",      "value" : 1    } ]  } ]}`

	_ = testInstantActions(t, vda5050s, json_str)

	return vda5050s

}

func Test_vda5050_InstantActions(t *testing.T) {

	vda5050s := vda5050_load_data(t)

	ret, err := vda5050s.State.Marshal()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	retStr := string(ret)

	if !strings.Contains(retStr, `[{"key":"nodeId","value":"6731b182088f262df9845fe2"},{"key":"isManual","value":false}`) {
		t.Error("testInstantActions updateLastNodeId actionType not found")
	}

	if !strings.Contains(retStr, `"actionType":"startPause"`) {
		t.Error("testInstantActions startPause actionType not found")

	}
	if !strings.Contains(retStr, `"actionType":"startJacking"`) {
		t.Error("testInstantActions startJacking actionType not found")
	}

}

func Test_vda5050_Action(t *testing.T) {

	vda5050s := vda5050_load_data(t)

	//vda5050s.State.ActionStates = []*state.ActionState{}

	actions := vda5050s.GetActions()

	for _, action := range actions {

		action.ActionStatus = state.Finished

		ret, err := json.Marshal(action)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(string(ret))

		break
	}

	s, err := vda5050s.GetStateJsonString()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(s))

	vda5050s.State.ActionStates = []*state.ActionState{}

	s, err = vda5050s.GetStateJsonString()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(s))

}

type MyLogger struct {
}

func NewMyLogger() *MyLogger {

	return &MyLogger{}
}

func (s *MyLogger) print(tag string, message string) {

	currentTime := time.Now().Format("2006-01-02 15:04:05.000")
	logMessage := fmt.Sprintf("[%s][%s]: %s", currentTime, tag, message)
	fmt.Println(logMessage)
}
func (s *MyLogger) INFO(message string) {
	s.print("INFO", message)
}

func (s *MyLogger) ERROR(message string) {
	s.print("ERROR", message)
}

func Test_vda5050_logger(t *testing.T) {

	vda5050s := vda5050_load_data(t)

	vda5050s.INFO("test1")

	vda5050s.SetLogger(NewMyLogger())

	vda5050s.INFO("test2")
	vda5050s.ERROR("test3")

}

func Benchmark_vda5050_Action(b *testing.B) {

}
