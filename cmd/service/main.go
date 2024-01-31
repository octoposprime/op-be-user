package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/golobby/container/v3"

	tconfig "github.com/octoposprime/op-be-shared/tool/config"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
	tredis "github.com/octoposprime/op-be-shared/tool/redis"
	pa_command "github.com/octoposprime/op-be-user/internal/application/presentation/adapter/command"
	pa_query "github.com/octoposprime/op-be-user/internal/application/presentation/adapter/query"
	as "github.com/octoposprime/op-be-user/internal/application/service"
	ds "github.com/octoposprime/op-be-user/internal/domain/service"
	ia_ebus "github.com/octoposprime/op-be-user/pkg/infrastructure/adapter/ebus"
	ia_repo "github.com/octoposprime/op-be-user/pkg/infrastructure/adapter/repository"
	ia_service "github.com/octoposprime/op-be-user/pkg/infrastructure/adapter/service"
	pc_grpc "github.com/octoposprime/op-be-user/pkg/presentation/controller/grpc"
	pc_probe "github.com/octoposprime/op-be-user/pkg/presentation/controller/probe"
	tseed "github.com/octoposprime/op-be-user/tool/config"
)

var internalConfig tconfig.InternalConfig
var dbConfig tconfig.DbConfig
var redisConfig tconfig.RedisConfig
var seedConfig tseed.SeedConfig

func main() {
	internalConfig.ReadConfig()
	dbConfig.ReadConfig()
	redisConfig.ReadConfig()
	seedConfig.ReadConfig()
	var err error

	fmt.Println("Starting User Service...")
	dbClient, err := tgorm.NewGormClient(tgorm.PostgresGormClient).Connect(dbConfig.PostgresDb.Host, dbConfig.PostgresDb.Port, dbConfig.PostgresDb.UserName, dbConfig.PostgresDb.Password, dbConfig.PostgresDb.Database)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")

	redisClient := tredis.NewRedisClient(redisConfig.Redis.Host, redisConfig.Redis.Port, redisConfig.Redis.Password, redisConfig.Redis.Db)
	_, err = redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis")

	cont := container.New()

	//Domain User Service
	err = cont.Singleton(func() *ds.Service {
		return ds.NewService()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure User Service Adapter
	err = cont.Singleton(func() ia_service.ServiceAdapter {
		return ia_service.NewServiceAdapter()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure User Db Repository Adapter
	err = cont.Singleton(func() ia_repo.DbAdapter {
		return ia_repo.NewDbAdapter(dbClient)
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure User Redis Repository Adapter
	err = cont.Singleton(func() ia_repo.RedisAdapter {
		return ia_repo.NewRedisAdapter(redisClient)
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure User EBus Adapter
	err = cont.Singleton(func() ia_ebus.EBusAdapter {
		return ia_ebus.NewEBusAdapter(redisClient)
	})
	if err != nil {
		panic(err)
	}

	//Application User Service
	err = cont.Singleton(func(s *ds.Service, d ia_repo.DbAdapter, r ia_repo.RedisAdapter, e ia_ebus.EBusAdapter, i ia_service.ServiceAdapter) *as.Service {
		return as.NewService(s, &d, &r, &e, &i)
	})
	if err != nil {
		panic(err)
	}

	//Application User Query Adapter
	err = cont.Singleton(func(s *as.Service) pa_query.QueryAdapter {
		return pa_query.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	//Application User Command Adapter
	err = cont.Singleton(func(s *as.Service) pa_command.CommandAdapter {
		return pa_command.NewCommandAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	var queryHandler pa_query.QueryAdapter
	err = cont.Resolve(&queryHandler)
	if err != nil {
		panic(err)
	}

	var commandHandler pa_command.CommandAdapter
	err = cont.Resolve(&commandHandler)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	if !internalConfig.Local {
		wg.Add(1)
		go pc_probe.NewProbeAPI().Serve(internalConfig.Restapi.ProbePort)
	}
	wg.Add(1)
	go pc_grpc.NewGrpc(queryHandler, commandHandler).Serve(internalConfig.Grpc.UserPort)
	wg.Wait()

}
