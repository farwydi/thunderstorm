package main

import (
	"github.com/farwydi/thunderstorm/domain"
	"github.com/farwydi/thunderstorm/usecase/command"
)

const (
	initial domain.State = iota
	start
	version
	gettingUserAge
	gettingUserEmail
	questionN1
	questionN2
)

func NewBasic() domain.TxTable {
	return basic
}

var basic = domain.TxTable{
	initial: {
		"command_start": domain.Route{
			NextState: start,
			Action:    command.Start,
		},
		"command_version": domain.Route{
			NextState: version,
			Action:    command.Start,
		},
	},

	version: {
		"command_start": domain.Route{
			NextState: start,
			Action:    command.Start,
		},
	},

	start: {
		"callback_collectingUserInformation": domain.Route{
			NextState: gettingUserAge,
			Action:    command.Start,
		},
		"callback_repeat": domain.Route{
			NextState: questionN1,
			Action:    command.Start,
		},
	},

	gettingUserAge: {
		"text": domain.Route{
			NextState: gettingUserEmail,
			Action:    command.Start,
		},
	},

	gettingUserEmail: {
		"text": domain.Route{
			NextState: questionN1,
			Action:    command.Start,
		},
	},

	questionN1: {
		"callback_answerN1": domain.Route{
			NextState: questionN2,
			Action:    command.Start,
		},
		"callback_answerN2": domain.Route{
			NextState: questionN2,
			Action:    command.Start,
		},
		"callback_answerN3": domain.Route{
			NextState: questionN2,
			Action:    command.Start,
		},
	},

	questionN2: {
		"callback_answerN1": domain.Route{
			NextState: start,
			Action:    command.Start,
		},
		"callback_answerN2": domain.Route{
			NextState: start,
			Action:    command.Start,
		},
		"callback_answerN3": domain.Route{
			NextState: start,
			Action:    command.Start,
		},
		"callback_back": domain.Route{
			NextState: questionN1,
			Action:    command.Start,
		},
	},
}
