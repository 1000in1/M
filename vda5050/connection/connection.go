// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    connection, err := UnmarshalConnection(bytes)
//    bytes, err = connection.Marshal()

package connection

import (
	"encoding/json"
	"time"
)

func UnmarshalConnection(data []byte) (Connection, error) {
	var r Connection
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Connection) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// The last will message of the AGV. Has to be sent with retain flag.
// Once the AGV comes online, it has to send this message on its connect topic, with the
// connectionState enum set to "ONLINE".
// The last will message is to be configured with the connection state set to
// "CONNECTIONBROKEN".
// Thus, if the AGV disconnects from the broker, master control gets notified via the topic
// "connection".
// If the AGV is disconnecting in an orderly fashion (e.g. shutting down, sleeping), the AGV
// is to publish a message on this topic with the connectionState set to "DISCONNECTED".
type Connection struct {
	// ONLINE: connection between AGV and broker is active. OFFLINE: connection between AGV and
	// broker has gone offline in a coordinated way. CONNECTIONBROKEN: The connection between
	// AGV and broker has unexpectedly ended.
	ConnectionState *ConnectionState `json:"connectionState,omitempty"`
	// Header ID of the message. The headerId is defined per topic and incremented by 1 with
	// each sent (but not necessarily received) message.
	HeaderID *int64 `json:"headerId,omitempty"`
	// Manufacturer of the AGV.
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Serial number of the AGV.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Timestamp in ISO8601 format (YYYY-MM-DDTHH:mm:ss.ssZ).
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Version of the protocol [Major].[Minor].[Patch]
	Version *string `json:"version,omitempty"`
}

// ONLINE: connection between AGV and broker is active. OFFLINE: connection between AGV and
// broker has gone offline in a coordinated way. CONNECTIONBROKEN: The connection between
// AGV and broker has unexpectedly ended.
type ConnectionState string

const (
	Connectionbroken ConnectionState = "CONNECTIONBROKEN"
	Offline          ConnectionState = "OFFLINE"
	Online           ConnectionState = "ONLINE"
)
