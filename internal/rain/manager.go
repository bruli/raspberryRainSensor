package rain

import "github.com/bruli/raspberryRainSensor/pkg/log"

//go:generate moq -out humidityReaderMock.go . HumidityReader
type HumidityReader interface {
	Read() (uint16, error)
}
type RainManager struct {
	reader HumidityReader
	Logger log.Logger
}

const rainRef = 300

func NewRainManager(reader HumidityReader, logger log.Logger) *RainManager {
	return &RainManager{reader: reader, Logger: logger}
}
func (m *RainManager) IsRaining() (bool, error) {
	v, err := m.readValues()
	if err != nil {
		return false, err
	}
	isRain := v < rainRef
	return isRain, err
}

func (m *RainManager) RainValue() (uint16, error) {
	v, err := m.readValues()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (m *RainManager) readValues() (uint16, error) {
	v, err := m.reader.Read()
	if err != nil {
		m.Logger.Fatalf("Fatal error reading humidity: %s", err)
	}
	return v, err
}
