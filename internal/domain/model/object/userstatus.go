package domain

// UserStatus is a status that represents the status of a user.
type UserStatus int8

const (
	UserStatusNONE UserStatus = iota
	UserStatusACTIVE
	UserStatusINACTIVE
)
