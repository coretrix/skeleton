package controller

import (
	hitrixController "github.com/coretrix/hitrix/pkg/controller"
	"github.com/gin-gonic/gin"
)

type DevPanelController struct {
}

func (controller *DevPanelController) GetActionListAction(c *gin.Context) {
	actions := []*hitrixController.MenuItem{
		{
			Label: "Clear Cache",
			URL:   "/dev/clear-cache/",
			Icon:  "mdiCached",
		},
	}

	c.JSON(200, actions)
}
