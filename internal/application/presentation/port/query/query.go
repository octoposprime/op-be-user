package application

// QueryPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the application layer.
type QueryPort interface {
	UserQueryPort
}
