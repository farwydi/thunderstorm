// +build wireinject

package main

import (
	"github.com/farwydi/thunderstorm/domain"
	"github.com/farwydi/thunderstorm/endpoint/webhook"
	"github.com/farwydi/thunderstorm/gateway/statedriver"
	"github.com/google/wire"
)

func setup(domain.Config) (application, func(), error) {
	panic(wire.Build(
		statedriver.NewStateDriverGateway,
		NewBasic,
		domain.NewStateMachine,
		webhook.NewRouter,
		newApplication,
	))
}
