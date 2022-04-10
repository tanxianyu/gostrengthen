//go:build wireinject
// +build wireinject

package main

import (
	"gostrengthen/week04/api/internel/biz"
	"gostrengthen/week04/api/internel/conf"
	"gostrengthen/week04/api/internel/server"
	"gostrengthen/week04/api/internel/service"

	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
)

type App struct {
	HttpServer *server.HttpServer
	GRPCServer *server.GRPCServer
	Client     *ent.Client
}

// newApp return App struct with server.HttpServer and server.GRPCServer
func newApp(http *server.HttpServer, grpc *server.GRPCServer, client *ent.Client) *App {
	return &App{HttpServer: http, GRPCServer: grpc, Client: client}
}

// initApp Inject wire ProvideSet
func initApp(group *errgroup.Group, option conf.Options) *App {
	panic(wire.Build(server.ProvideSet, data.ProvideSet, service.ProvideSet, biz.ProvideSet, newApp))
}
