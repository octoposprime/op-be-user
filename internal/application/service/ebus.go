package application

import (
	"context"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	me "github.com/octoposprime/op-be-user/internal/domain/model/entity"
)

// This is the event listener handler of the application layer.
func (a *Service) EventListen() *Service {
	go a.Listen(context.Background(), smodel.ChannelCreateUser, a.EventListenerCallBack)
	go a.Listen(context.Background(), smodel.ChannelDeleteUser, a.EventListenerCallBack)
	return a
}

// This is a call-back function of the event listener handler of the application layer.
func (a *Service) EventListenerCallBack(channelName string, user me.User) {
	if channelName == smodel.ChannelCreateUser {
		// Used a.CreateUser instead of this method for Redis Caching
		//a.SaveUser(context.Background(), user)
		a.CreateUser(context.Background(), user)
	} else if channelName == smodel.ChannelDeleteUser {
		a.DeleteUser(context.Background(), user)
	} else {
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "EventListenerCallBack", channelName, smodel.ErrorChannelNameNotValid.Error()))
	}
}
