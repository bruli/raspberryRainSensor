package rain

//go:generate moq -out rainRepositoryMock.go . RainRepository
type RainRepository interface {
	Read() (uint16, error)
}
