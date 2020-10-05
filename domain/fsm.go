package domain

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	State int

	Action func(ctx context.Context, currentState, nextState State, routeKey string, update tgbotapi.Update) error

	Route struct {
		NextState State
		Action    Action
	}

	TxTable map[State]map[string]Route
)

func NewStateMachine(sd StateDriverGateway, ttx TxTable) *StateMachine {
	return &StateMachine{
		sd:               sd,
		transactionTable: ttx,
	}
}

type StateMachine struct {
	sd               StateDriverGateway
	transactionTable TxTable
}

func (m *StateMachine) Tx(ctx context.Context, update tgbotapi.Update) error {
	routerKey := ""
	var user *tgbotapi.User
	if update.CallbackQuery != nil {
		routerKey = "callback_" + update.CallbackQuery.Data
		user = update.CallbackQuery.Message.From
	}
	if update.Message != nil {
		routerKey = "text"
		user = update.Message.From

		if update.Message.IsCommand() {
			routerKey = "command_" + update.Message.Command()
		}
	}

	currentState, err := m.sd.CurrentState(ctx, user)
	if err != nil {
		return err
	}

	table := m.transactionTable[currentState]
	route, found := table[routerKey]
	if !found {
		return fmt.Errorf("no tx")
	}

	return m.sd.PushTx(ctx, user, route.NextState, func() error {
		return route.Action(ctx, currentState, route.NextState, routerKey, update)
	})
}
