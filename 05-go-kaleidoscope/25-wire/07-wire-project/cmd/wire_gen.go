// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"go-micro-frame-doc/25-wire/07-wire-project/internal/data"
	"go-micro-frame-doc/25-wire/07-wire-project/internal/server/config"
	"go-micro-frame-doc/25-wire/07-wire-project/internal/server/db"
)

// Injectors from wire.go:

//go:generate wire
func InitApp(ctx context.Context) (*App, func(), error) {
	configConfig, cleanup, err := config.New(ctx)
	if err != nil {
		return nil, nil, err
	}
	sqlDB, cleanup2, err := db.NewDb(ctx, configConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dataData, cleanup3, err := data.NewData(sqlDB)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	orderRepo, cleanup4, err := data.NewOrderRepo(dataData)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := NewApp(configConfig, sqlDB, orderRepo)
	return app, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
