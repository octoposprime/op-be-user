package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorUserNotFound,
	ErrorUserIdIsEmpty,
	ErrorUserIsEmpty,
	ErrorUserIsInactive,
	ErrorUserUsernameOrEmailAreEmpty,
	ErrorUserTokenIsEmpty,
	ErrorUserAuthenticationTokenAndRefreshTokenAreEmpty,
	ErrorUserUsernameIsEmpty,
	ErrorUserUsernameIsExists,
	ErrorUserUsernameIsTooShort,
	ErrorUserUsernameIsTooLong,
	ErrorUserUsernameContainsInValidChars,
	ErrorUserUsernameContainsSpecialChars,
	ErrorUserUsernameIsNotValid,
	ErrorUserEmailIsEmpty,
	ErrorUserEmailIsNotValid,
	ErrorUserEmailIsExists,
	ErrorUserRoleIsEmpty,
	ErrorUserPasswordIsEmpty,
	ErrorUserPasswordIsTooShort,
	ErrorUserPasswordIsTooLong,
	ErrorUserPasswordIsNotValid,
	ErrorUserPasswordNotFound,
	ErrorUserPasswordIsInactive,
}

const (
	ErrId              string = "id"
	ErrUser            string = "user"
	ErrUserName        string = "username"
	ErrPassword        string = "password"
	ErrToken           string = "token"
	ErrAuthentication  string = "authentication"
	ErrRefrehsToken    string = "refreshtoken"
	ErrEmail           string = "email"
	ErrRole            string = "role"
	ErrUserNameOrEmail string = "username_or_email"
)

const (
	ErrEmpty         string = "empty"
	ErrTooShort      string = "tooshort"
	ErrTooLong       string = "toolong"
	ErrContainsInvalidChars string = "containsinvalidchars"
	ErrContainsSpecialChars string = "containsspecialchars"
	ErrNotValid      string = "notvalid"
	ErrInactive      string = "inactive"
	ErrAlreadyExists string = "alreadyexists"
)

var (
	ErrorNone                                           error = nil
	ErrorUserNotFound                                   error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + smodel.ErrNotFound)
	ErrorUserIdIsEmpty                                  error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrId + smodel.ErrSep + ErrEmpty)
	ErrorUserIsEmpty                                    error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrEmpty)
	ErrorUserIsInactive                                 error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrInactive)
	ErrorUserUsernameOrEmailAreEmpty                    error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserNameOrEmail + smodel.ErrSep + ErrEmpty)
	ErrorUserTokenIsEmpty                               error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrToken + smodel.ErrSep + ErrEmpty)
	ErrorUserAuthenticationTokenAndRefreshTokenAreEmpty error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrAuthentication + smodel.ErrSep + ErrRefrehsToken + smodel.ErrSep + ErrEmpty)
	ErrorUserUsernameIsEmpty                            error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserName + smodel.ErrSep + ErrEmpty)
	ErrorUserUsernameIsExists                           error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserName + smodel.ErrSep + ErrAlreadyExists)
	ErrorUserUsernameIsTooShort                         error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserName + smodel.ErrSep + ErrTooShort)
	ErrorUserUsernameIsTooLong                          error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserName + smodel.ErrSep + ErrTooLong)
	ErrorUserUsernameContainsInValidChars               error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserName + smodel.ErrSep + ErrContainsInvalidChars)
	ErrorUserUsernameContainsSpecialChars               error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserName + smodel.ErrSep + ErrorUserUsernameContainsSpecialChars )         
	ErrorUserUsernameIsNotValid                         error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrUserName + smodel.ErrSep + ErrNotValid)
	ErrorUserEmailIsEmpty                               error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrEmail + smodel.ErrSep + ErrEmpty)
	ErrorUserEmailIsNotValid                            error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrEmail + smodel.ErrSep + ErrNotValid)
	ErrorUserEmailIsExists                              error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrEmail + smodel.ErrSep + ErrAlreadyExists)
	ErrorUserRoleIsEmpty                                error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrRole + smodel.ErrSep + ErrEmpty)

	ErrorUserPasswordIsEmpty    error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrPassword + smodel.ErrSep + ErrEmpty)
	ErrorUserPasswordIsTooShort error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrPassword + smodel.ErrSep + ErrTooShort)
	ErrorUserPasswordIsTooLong  error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrPassword + smodel.ErrSep + ErrTooLong)
	ErrorUserPasswordIsNotValid error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrPassword + smodel.ErrSep + ErrNotValid)
	ErrorUserPasswordNotFound   error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrPassword + smodel.ErrSep + smodel.ErrNotFound)
	ErrorUserPasswordIsInactive error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrUser + smodel.ErrSep + ErrPassword + smodel.ErrSep + ErrInactive)
)

func GetErrors() []error {
	return ERRORS
}
