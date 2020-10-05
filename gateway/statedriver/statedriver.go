package statedriver

import (
	"context"
	"github.com/farwydi/thunderstorm/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewStateDriverGateway() domain.StateDriverGateway {
	return &stateDriverGateway{}
}

type stateDriverGateway struct {
}

func (s *stateDriverGateway) PushTx(ctx context.Context, user *tgbotapi.User, nextState domain.State, tx func() error) error {
	panic("implement me")
}

func (s *stateDriverGateway) CurrentState(ctx context.Context, user *tgbotapi.User) (domain.State, error) {
	panic("implement me")
}
