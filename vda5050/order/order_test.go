package order

import (
	"fmt"
	"testing"
)

func Test_order_unmarshal(t *testing.T) {

	var json_str = `{  "headerId" : 1903,  "timestamp" : "2024-11-29T05:23:29.044972Z",  "version" : "2.0.0",  "manufacturer" : "test",  "serialNumber" : "test",  "orderId" : "Park-01JDV5ZHY0ZK7Y8JE59JZ9F1JN-0",  "orderUpdateId" : 0,  "nodes" : [ {    "nodeId" : "n_22b25ddc4dcd7a2bd822d1e2e7f82470",    "sequenceId" : 0,    "released" : true,    "actions" : [ ],    "nodePosition" : {      "x" : 18.897,      "y" : -29.844,      "mapId" : "673d98a136bbc50b54d4381e",      "allowedDeviationXY" : 0.3638997033058943,      "allowedDeviationTheta" : 3.141592653589793    }  }, {    "nodeId" : "673d98d4f2d42efb5966d862",    "sequenceId" : 2,    "released" : true,    "actions" : [ ],    "nodePosition" : {      "x" : 17.869,      "y" : -28.355,      "mapId" : "673d98a136bbc50b54d4381e",      "theta" : 2.349038576343302,      "allowedDeviationXY" : 0.3,      "allowedDeviationTheta" : 3.141592653589793    }  } ],  "edges" : [ {    "edgeId" : "l_8c1720d9eb300cf151d387799f9a5e55",    "sequenceId" : 1,    "released" : true,    "startNodeId" : "n_22b25ddc4dcd7a2bd822d1e2e7f82470",    "endNodeId" : "673d98d4f2d42efb5966d862",    "actions" : [ ],    "maxSpeed" : 1.0  } ],  "routeIndex" : 0,  "routePoints" : [ "n_22b25ddc4dcd7a2bd822d1e2e7f82470", "673d98d4f2d42efb5966d862" ],  "routePaths" : [ {    "x" : 18897,    "y" : -29844,    "z" : 0  }, {    "x" : 17869,    "y" : -28355,    "z" : 0  } ],  "destinationName" : "673d98d4f2d42efb5966d862",  "allocatedResources" : [ "n_22b25ddc4dcd7a2bd822d1e2e7f82470", "673d98d4f2d42efb5966d862" ],  "finalMovement" : true,  "finalOrder" : true,  "taskId" : "Park-01JDV5ZHY0ZK7Y8JE59JZ9F1JN"}`

	vda5050_order, err := UnmarshalOrder([]byte(json_str))
	if err != nil {
		t.Log("UnmarshalOrder error", err)
	}
	fmt.Println(*vda5050_order.OrderID, *vda5050_order.OrderUpdateID)

	for _, node := range vda5050_order.Nodes {
		fmt.Println(*node.NodeID, *node.NodePosition.X, *node.NodePosition.Y)
	}

	fmt.Println(vda5050_order.Nodes)

}
