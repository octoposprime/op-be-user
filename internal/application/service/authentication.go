package application

import (
	"context"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tjwt "github.com/octoposprime/op-be-shared/tool/jwt"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-user/internal/domain/model/object"
)

// Login generates an authentication token if the given login request values are valid.
func (a *Service) Login(ctx context.Context, loginRequest mo.LoginRequest) (mo.Token, error) {
	if err := loginRequest.Validate(); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	var userFilter me.UserFilter
	if loginRequest.UserName != "" {
		userFilter.UserName = loginRequest.UserName
	} else {
		userFilter.Email = loginRequest.Email
	}
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	if users.TotalRows > 0 {
		inUser := users.Users[0]
		if err := a.CheckIsAuthenticable(&inUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
			return *mo.NewEmptyToken(), err
		}
		userPassword, err := a.GetUserPasswordByUserId(ctx, inUser.Id)
		if err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
			return *mo.NewEmptyToken(), err
		}
		if userPassword.PasswordStatus == mo.PasswordStatusINACTIVE {
			err = mo.ErrorUserPasswordIsInactive
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
			return *mo.NewEmptyToken(), err
		}
		if userPassword.ComparePass(loginRequest.Password) {
			userInfo := inUser
			claims := tjwt.NewClaims(userInfo.Id.String(), userInfo)
			accessToken, refreshToken, err := claims.GenerateJWT()
			if err != nil {
				userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
				go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Login", userId, err.Error()))
				return *mo.NewEmptyToken(), err
			}
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeINFO, "Login", userId, "login succeeded"))
			return *mo.NewToken(accessToken, refreshToken), nil
		}
		return *mo.NewEmptyToken(), smodel.ErrorUserNotFound
	}
	return *mo.NewEmptyToken(), smodel.ErrorUserNotFound
}

// Refresh regenerate an authentication token.
func (a *Service) Refresh(ctx context.Context, token mo.Token) (mo.Token, error) {
	if err := token.Validate(); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Refresh", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	userId, err := tjwt.DecodeRefreshJWT(token.RefreshToken)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Refresh", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	var userFilter me.UserFilter
	userFilter.Id = tuuid.FromString(userId)
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Refresh", userId, err.Error()))
		return *mo.NewEmptyToken(), err
	}
	if users.TotalRows > 0 {
		inUser := users.Users[0]
		userInfo := inUser
		if err := a.CheckIsAuthenticable(&inUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Refresh", userId, err.Error()))
			return *mo.NewEmptyToken(), err
		}
		claims := tjwt.NewClaims(userInfo.Id.String(), userInfo)
		accessToken, refreshToken, err := claims.GenerateJWT()
		if err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Refresh", userId, err.Error()))
			return *mo.NewEmptyToken(), err
		}
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeINFO, "Refresh", userId, "refresh succeeded"))
		return *mo.NewToken(accessToken, refreshToken), nil
	}
	return *mo.NewEmptyToken(), smodel.ErrorUserNotFound
}

// Logout clears some footprints for the user.
func (a *Service) Logout(ctx context.Context, token mo.Token) error {
	if err := token.Validate(); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Logout", userId, err.Error()))
		return err
	}
	userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "Logout", userId, "logout succeeded"))
	return nil
}
