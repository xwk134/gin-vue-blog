package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(map[string]string{
		"id": "001",
	}, c)
}
