package mocks

import (
	hymnIOC "coretrix/skeleton/pkg/ioc"
	"coretrix/skeleton/pkg/ioc/service/oss"

	"github.com/coretrix/hitrix/service"

	"github.com/sarulabs/di"
)

func FakeOSSService(fakeOSSService oss.Client) *service.Definition {
	return &service.Definition{
		Name:   hymnIOC.OSSService,
		Global: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &fakeOSSService, nil
		},
	}
}
