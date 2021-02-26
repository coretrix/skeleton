package main

import (
	"coretrix/skeleton/pkg/entity"
	"coretrix/skeleton/pkg/script"

	hitrixRegistry "github.com/coretrix/hitrix/service/registry"

	"github.com/coretrix/hitrix"
)

func main() {
	h, deferFunc := hitrix.New(
		"subscribers-notifier", "secret",
	).RegisterDIService(
		hitrixRegistry.ServiceProviderConfigDirectory("../../config"),
		hitrixRegistry.ServiceDefinitionSlackAPI(),
		hitrixRegistry.ServiceProviderErrorLogger(),
		hitrixRegistry.ServiceDefinitionOrmRegistry(entity.Init),
		hitrixRegistry.ServiceDefinitionOrmEngine(),
	).Build()
	defer deferFunc()

	h.RunScript(&script.SubscriberNotifier{})
}
