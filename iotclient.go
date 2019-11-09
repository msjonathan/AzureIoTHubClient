package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/amenzhinsky/iothub/iotdevice"
	iotmqtt "github.com/amenzhinsky/iothub/iotdevice/transport/mqtt"
	"github.com/tkanos/gonfig"
)

// deviceStatusChanged message
type deviceStatusChanged struct {
	Body string
}

// Configuration is representation of the json config
type Configuration struct {
	IotHubDeviceConnectionString string
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("./config.json", &configuration)

	c, err := iotdevice.NewFromConnectionString(
		iotmqtt.New(), configuration.IotHubDeviceConnectionString,
	)
	if err != nil {
		log.Fatal(err)
	}

	// connect to the iothub
	if err = c.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("Welcome %d times\n", i)

		data := struct{ Temperature int }{rand.Intn(35)}
		body, err := json.Marshal(data)

		deviceStatusChangedMessage := deviceStatusChanged{
			Body: string(body)}

		message, err := json.Marshal(deviceStatusChangedMessage)

		log.Printf("Sending: %s", string(message))

		if err != nil {
			log.Fatal(err)
		}

		// send a device-to-cloud message
		if err = c.SendEvent(context.Background(), message, iotdevice.WithSendProperty("Type", "DeviceStatusChanged")); err != nil {
			log.Fatal(err)
		}

		time.Sleep(250 * time.Microsecond)
	}

}
