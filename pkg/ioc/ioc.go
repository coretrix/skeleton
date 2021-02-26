package ioc

import (
	"context"
	"coretrix/skeleton/pkg/ioc/service/oss"

	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/component/config"
	slackapi "github.com/coretrix/hitrix/service/component/slack_api"

	"github.com/summer-solutions/orm"
)

const (
	OSSService = "oss_service"
)

func GetOSSService() oss.Client {
	return service.GetServiceRequired(OSSService).(oss.Client)
}

func GetSlackAPIService() *slackapi.SlackAPI {
	return service.GetServiceRequired("slack_api").(*slackapi.SlackAPI)
}
func GetConfigService() *config.Config {
	return service.GetServiceRequired("config").(*config.Config)
}

func GetOrmEngineGlobalService() *orm.Engine {
	return service.GetServiceRequired("orm_engine_global").(*orm.Engine)
}

func GetOrmEngineFromContext(ctx context.Context) *orm.Engine {
	return service.GetServiceForRequestRequired(ctx, "orm_engine_request").(*orm.Engine)
}
