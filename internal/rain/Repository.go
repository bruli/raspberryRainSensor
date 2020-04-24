package rain

//go:generate moq -out RepositoryMock.go . Repository
type Repository interface {
	Read() (uint16, error)
}
