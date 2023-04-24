package hunter

import (
	"context"

	"github.com/bruli/raspberryRainSensor/internal/domain/rain"
	"github.com/stianeikeland/go-rpio/v4"
)

const RainReference = 600

type RainSensor struct{}

func (r RainSensor) Read(ctx context.Context) (rain.Rain, error) {
	select {
	case <-ctx.Done():
		return rain.Rain{}, ctx.Err()
	default:
		if err := rpio.Open(); err != nil {
			return rain.Rain{}, err
		}

		defer func() {
			_ = rpio.Close()
		}()

		if err := rpio.SpiBegin(rpio.Spi0); err != nil {
			return rain.Rain{}, err
		}

		rpio.SpiSpeed(1000000)
		rpio.SpiChipSelect(0)
		channel := byte(0)
		data := []byte{1, (8 + channel) << 4, 0}

		rpio.SpiExchange(data)

		value := int(data[1]&3)<<8 + int(data[2])
		defer rpio.SpiEnd(rpio.Spi0)

		return rain.New(value, value > RainReference), nil
	}
}
