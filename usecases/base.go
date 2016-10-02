package usecases

// Logger interface
type Logger interface {
	Log(message interface{}) error
}
