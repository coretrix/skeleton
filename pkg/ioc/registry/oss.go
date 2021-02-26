package registry

import (
	"coretrix/skeleton/pkg/ioc"
	"coretrix/skeleton/pkg/ioc/service/oss/storage"

	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/component/app"
	"github.com/coretrix/hitrix/service/component/config"

	"github.com/sarulabs/di"
)

func OSS() *service.Definition {
	return &service.Definition{
		Name:   ioc.OSSService,
		Global: true,
		Build: func(ctn di.Container) (interface{}, error) {
			configService := ctn.Get("config").(*config.Config)
			appService := ctn.Get("app").(*app.App)
			return storage.NewGoogleOSS(configService.GetFolderPath()+"/.oss.json", appService.Mode), nil
		},
	}
}
