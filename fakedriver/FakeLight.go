package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ninjasphere/go-ninja/api"
	"github.com/ninjasphere/go-ninja/channels"
	"github.com/ninjasphere/go-ninja/model"
)

type FakeLight struct {
	driver       ninja.Driver
	info         *model.Device
	sendEvent    func(event string, payload interface{}) error
	onOffChannel *channels.OnOffChannel
}

func NewFakeLight(driver ninja.Driver, id int) *FakeLight {
	name := fmt.Sprintf("Fancy Fake Light %d", id)
	
	light := &FakeLight{
		driver: driver,
		info: &model.Device{
			NaturalID:     fmt.Sprintf("light%d", id),
			NaturalIDType: "fake",
			Name:          &name,
			Signatures: &map[string]string{
				"ninja:manufacturer": "Fake Co.",
				"ninja:productName":  "FakeLight",
				"ninja:productType":  "Light",
				"ninja:thingType":    "light",
			},
		},
	}

	light.onOffChannel = channels.NewOnOffChannel(light)

	return light
}

func (l *FakeLight) GetDeviceInfo() *model.Device {
	return l.info
}

func (l *FakeLight) GetDriver() ninja.Driver {
	return l.driver
}

func (l *FakeLight) SetOnOff(state bool) error {
	log.Printf("Turning %t", state)
	return nil
}

func (l *FakeLight) ToggleOnOff() error {
	log.Println("Toggling")
	return nil
}

func (l *FakeLight) SetEventHandler(sendEvent func(event string, payload interface{}) error) {
	l.sendEvent = sendEvent
}

var reg, _ = regexp.Compile("[^a-z0-9]")

// Exported by service/device schema
func (l *FakeLight) SetName(name *string) (*string, error) {
	log.Printf("Setting device name to %s", *name)

	safe := reg.ReplaceAllString(strings.ToLower(*name), "")
	if len(safe) > 5 {
		safe = safe[0:5]
	}

	log.Printf("Pretending we can only set 5 lowercase alphanum. Name now: %s", safe)

	l.sendEvent("renamed", safe)

	return &safe, nil
}