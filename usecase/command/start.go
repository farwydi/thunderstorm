package command

import (
	"context"
	"github.com/farwydi/thunderstorm/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(ctx context.Context, currentState, nextState domain.State, routeKey string, update tgbotapi.Update) error {
	user := update.Message.From

	return nil
}
