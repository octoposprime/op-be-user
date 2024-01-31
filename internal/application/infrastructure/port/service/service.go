package application

// ServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other servies.
type ServicePort interface {
	LoggingServicePort
}
