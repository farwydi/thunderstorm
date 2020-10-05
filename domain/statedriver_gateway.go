package domain

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StateDriverGateway interface {
	PushTx(ctx context.Context, user *tgbotapi.User, nextState State, tx func() error) error
	CurrentState(ctx context.Context, user *tgbotapi.User) (State, error)
}
