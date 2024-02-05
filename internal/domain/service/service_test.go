package domain

import (
	"testing"

	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
)

func TestService_CheckUserNameRules(t *testing.T) {
	type args struct {
		user *me.User
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if err := s.CheckUserNameRules(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Service.CheckUserNameRules() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
