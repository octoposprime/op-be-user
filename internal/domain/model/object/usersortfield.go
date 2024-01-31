package domain

// UserSortField is a type that represents the sort fields of a user.
type UserSortField int8

const (
	UserSortFieldNONE UserSortField = iota
	UserSortFieldId
	UserSortFieldCompanyId
	UserSortFieldName
	UserSortFieldCreatedAt
	UserSortFieldUpdatedAt
)
