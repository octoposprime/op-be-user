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
		{
			name: "Empty username",
			s:    NewService(),
			args: args{
				user: &me.User{
					UserName: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Username is too short",
			s:    NewService(),
			args: args{
				user: &me.User{
					UserName: "short",
				},
			},
			wantErr: true,
		},
		{
			name: "Username is too long",
			s:    NewService(),
			args: args{
				user: &me.User{
					UserName: "thisusernameistoolong",
				},
			},
			wantErr: true,
		},
		{
			name: "Username contains space",
			s:    NewService(),
			args: args{
				user: &me.User{
					UserName: "user name",
				},
			},
			wantErr: true,
		},
		{
			name: "Username contains special characters",
			s:    NewService(),
			args: args{
				user: &me.User{
					UserName: "asd!^+.",
				},
			},
			wantErr: true,
		},
		{
			name: "Username is valid",
			s:    NewService(),
			args: args{
				user: &me.User{
					UserName: "Qwe123_",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.CheckUserNameRules(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Service.CheckUserNameRules() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
