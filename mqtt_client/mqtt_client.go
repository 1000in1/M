package mqttclient

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type AxMqttClient struct {
	client_id      string
	mqtt_broker    string
	mqtt_user      string
	mqtt_pass      string
	client         mqtt.Client
	onConnectFunc  func(client AxMqttClient)
	onMessagetFunc func(topic string, pyload []byte)
}

func NewAxMqttClient(client_id string, mqtt_broker string, mqtt_user string, mqtt_pass string) *AxMqttClient {
	return &AxMqttClient{
		client_id:   client_id,
		mqtt_broker: mqtt_broker,
		mqtt_user:   mqtt_user,
		mqtt_pass:   mqtt_pass,
	}
}

func (t *AxMqttClient) Disconnect() {
	t.client.Disconnect(250)
}

func (t *AxMqttClient) Subscribe(topic string, qos byte) bool {

	if token := t.client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		return false
	}
	return true
}
func (t *AxMqttClient) Connect(onConnectFunc func(client AxMqttClient), onMessagetFunc func(topic string, pyload []byte)) bool {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(t.mqtt_broker)
	opts.SetClientID(t.client_id)
	opts.SetUsername(t.mqtt_user)
	opts.SetPassword(t.mqtt_pass)
	opts.SetCleanSession(true)
	opts.SetKeepAlive(15 * time.Second)
	opts.SetDefaultPublishHandler(t.messageHandler)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetConnectRetry(true)
	opts.SetConnectionLostHandler(t.connectionLostHandler)
	opts.SetOnConnectHandler(t.onConnect)

	t.client = mqtt.NewClient(opts)

	t.onConnectFunc = onConnectFunc
	t.onMessagetFunc = onMessagetFunc

	// Connect to the broker
	if token := t.client.Connect(); token.Wait() && token.Error() != nil {
		return false
	}

	return true

}

func (t *AxMqttClient) messageHandler(client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	t.onMessagetFunc(msg.Topic(), msg.Payload())
}

// connectionLostHandler handles lost connection and attempts to reconnect
func (t *AxMqttClient) connectionLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)

}

func (t *AxMqttClient) onConnect(client mqtt.Client) {
	fmt.Printf("onConnect \n")
	t.onConnectFunc(*t)

}
