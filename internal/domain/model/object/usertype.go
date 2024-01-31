package domain

// UserType is a type that represents the type of a user.
type UserType int8

const (
	UserTypeNONE UserType = iota
	UserTypeADMIN
	UserTypeUSER
	UserTypePARTNER
	UserTypeANALYST
	UserTypePM
)
