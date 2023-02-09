package apis

import (
	"github.com/LINQQ1212/common2/global"
	"github.com/LINQQ1212/common2/response"
	"github.com/gin-gonic/gin"
)

func VersionInfo(c *gin.Context) {
	v := c.Query("v")
	version, ok := global.Versions.Get(v)
	if !ok {
		response.FailWithMessage("版本不存在", c)
		return
	}
	data := map[string]any{}
	data["info"] = version.Info
	data["domains"] = version.Domains
	data["reviewNum"] = version.ReviewNum
	data["categoriesLen"] = version.CategoriesLen
	response.OkWithData(data, c)

}
