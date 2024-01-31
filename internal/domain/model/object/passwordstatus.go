package domain

// PasswordStatus is a type that represents the status of a password.
type PasswordStatus int8

const (
	PasswordStatusNONE PasswordStatus = iota
	PasswordStatusACTIVE
	PasswordStatusINACTIVE
	PasswordStatusAUTO_GENERATED
	PasswordStatusCHANGE_REQUIRED
	PasswordStatusEXPIRED
)
