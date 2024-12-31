package mqttclient

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/1000in1/m/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	tag            string
	client_id      string
	mqtt_broker    string
	mqtt_user      string
	mqtt_pass      string
	logger         logger.LoggerIF
	client         mqtt.Client
	onConnectFunc  func(client MqttClient)
	onMessagetFunc func(topic string, pyload []byte)
}

func NewMqttClient(client_id string, mqtt_broker string, mqtt_user string, mqtt_pass string) *MqttClient {
	return &MqttClient{
		tag:         "AxMqttClient",
		client_id:   client_id,
		mqtt_broker: mqtt_broker,
		mqtt_user:   mqtt_user,
		mqtt_pass:   mqtt_pass,
		logger:      nil,
	}
}

func (t *MqttClient) Disconnect() {
	t.client.Disconnect(250)
}

func (t *MqttClient) INFO(message string) {
	if t.logger != nil {
		t.logger.INFO(t.tag, message)
	}
}

func (t *MqttClient) ERROR(message string) {
	if t.logger != nil {
		t.logger.ERROR(t.tag, message)
	}
}

func (t *MqttClient) SetLogger(logger *logger.Logger) {
	t.logger = logger
}

func (t *MqttClient) Subscribe(topic string, qos byte) bool {

	if token := t.client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		return false
	}
	return true
}
func (t *MqttClient) Publish(topic string, pyload []byte) {

	qos := byte(0) // 消息质量等级
	t.INFO(fmt.Sprintf("Publish: %s -> %s", topic, string(pyload)))
	t.client.Publish(topic, qos, false, pyload)
}

func (t *MqttClient) PublishEx(topic string, pyload []byte, qos byte, retained bool) {

	t.client.Publish(topic, qos, false, pyload)
}

func (t *MqttClient) Connect(onConnectFunc func(client MqttClient), onMessagetFunc func(topic string, pyload []byte)) bool {

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

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	opts.SetTLSConfig(tlsConfig)

	t.client = mqtt.NewClient(opts)

	t.onConnectFunc = onConnectFunc
	t.onMessagetFunc = onMessagetFunc

	t.INFO("Connecting to broker")

	// Connect to the broker
	if token := t.client.Connect(); token.Wait() && token.Error() != nil {
		t.INFO("Connecting to error")
		return false
	}
	t.INFO("Connecting to broker")
	return true

}

func (t *MqttClient) messageHandler(client mqtt.Client, msg mqtt.Message) {
	//fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	t.onMessagetFunc(msg.Topic(), msg.Payload())
}

// connectionLostHandler handles lost connection and attempts to reconnect
func (t *MqttClient) connectionLostHandler(client mqtt.Client, err error) {

	t.ERROR(fmt.Sprintf("Connection lost: %v\n", err))

}

func (t *MqttClient) onConnect(client mqtt.Client) {
	t.INFO("onConnect")
	t.onConnectFunc(*t)

}
