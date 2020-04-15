package humiditySensor

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type Sensor struct {
}

func NewSensor() *Sensor {
	return &Sensor{}
}

func (s *Sensor) Read() (uint16, error) {
	if err := rpio.Open(); err != nil {
		return 0, err
	}

	defer rpio.Close()

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		return 0, err
	}

	rpio.SpiSpeed(1000000)
	rpio.SpiChipSelect(0)
	channel := byte(0)
	data := []byte{1, (8 + channel) << 4, 0}

	rpio.SpiExchange(data)

	code := uint16(data[1]&3)<<8 + uint16(data[2])
	defer rpio.SpiEnd(rpio.Spi0)

	return code, nil
}
