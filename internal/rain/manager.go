package rain

import "github.com/bruli/raspberryRainSensor/pkg/log"

//go:generate moq -out humidityReaderMock.go . HumidityReader
type HumidityReader interface {
	Read() (uint16, error)
}
type Manager struct {
	reader HumidityReader
	Logger log.Logger
}

const rainRef = 300

func NewManager(reader HumidityReader, logger log.Logger) *Manager {
	return &Manager{reader: reader, Logger: logger}
}
func (m *Manager) IsRaining() (bool, error) {
	v, err := m.reader.Read()
	if err != nil {
		m.Logger.Fatalf("Fatal error reading humidity: %s", err)
		return false, err
	}

	isRain := v < rainRef
	return isRain, err
}
