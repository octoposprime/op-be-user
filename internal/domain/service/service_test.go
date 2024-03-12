package domain_test

import (
    "testing"

    me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
    "github.com/octoposprime/op-be-user/internal/domain"
    "github.com/stretchr/testify/require"
)

func TestService_CheckUserNameRules(t *testing.T) {
    service := domain.NewService()

    tests := []struct {
        name      string
        userName  string
        expectErr bool
    }{
        {"containsSpecialChars", "asd!^+.", true},
        {"noSpecialChars", "Qwe123_", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            user := &me.User{UserName: tt.userName}
            err := service.CheckUserNameRules(user)
            if tt.expectErr {
                require.Error(t, err, "expected an error for username: %s", tt.userName)
            } else {
                require.NoError(t, err, "did not expect an error for username: %s", tt.userName)
            }
        })
    }
}
