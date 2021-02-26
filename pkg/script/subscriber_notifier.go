package script

import (
	"context"

	"github.com/coretrix/hitrix"
)

type SubscriberNotifier struct {
}

func (script *SubscriberNotifier) Run(_ context.Context, _ hitrix.Exit) {
	//todo your logic here
}

func (script *SubscriberNotifier) Unique() bool {
	return true
}

func (script *SubscriberNotifier) Code() string {
	return "test-code"
}

func (script *SubscriberNotifier) Description() string {
	return "Test code"
}
