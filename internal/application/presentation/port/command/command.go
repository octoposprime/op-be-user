package application

// CommandPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type CommandPort interface {
	LoggingCommandPort
	AuthenticationCommandPort
	UserCommandPort
}
