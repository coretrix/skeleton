package main

import (
	"coretrix/skeleton/api/web-api/graphql/graph"
	"coretrix/skeleton/api/web-api/graphql/graph/generated"
	"coretrix/skeleton/api/web-api/rest/middleware"
	"coretrix/skeleton/pkg/entity"
	"coretrix/skeleton/pkg/ioc/registry"

	"github.com/coretrix/hitrix"
	hitrixRegistry "github.com/coretrix/hitrix/service/registry"

	hitrixMiddleware "github.com/coretrix/hitrix/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	s, deferFunc := hitrix.New(
		"web-api", "secret",
	).RegisterDIService(
		hitrixRegistry.ServiceProviderConfigDirectory("../../config"),
		hitrixRegistry.ServiceDefinitionSlackAPI(),
		hitrixRegistry.ServiceProviderErrorLogger(),
		hitrixRegistry.ServiceDefinitionOrmRegistry(entity.Init),
		hitrixRegistry.ServiceDefinitionOrmEngine(), //web need that for tests
		hitrixRegistry.ServiceDefinitionOrmEngineForContext(),
		hitrixRegistry.ServiceProviderJWT(),
		hitrixRegistry.ServiceProviderPassword(),
		registry.OSS(),
	).
		RegisterDevPanel(&entity.AdminUserEntity{}, hitrixMiddleware.Router, nil).Build()
	defer deferFunc()

	config := generated.Config{Resolvers: &graph.Resolver{}, Directives: generated.DirectiveRoot{Validate: hitrix.ValidateDirective()}}

	s.RunServer(4001, generated.NewExecutableSchema(config), func(ginEngine *gin.Engine) {
		hitrixMiddleware.Cors(ginEngine)
		middleware.Router(ginEngine)
	})
}
