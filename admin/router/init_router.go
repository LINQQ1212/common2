package router

import (
	"github.com/LINQQ1212/common2/admin/apis"
	"github.com/LINQQ1212/common2/middleware"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.GET("/version.txt", apis.VersionPHP)
	r.GET("/version-jm.txt", apis.VersionPHP)
	r.GET("/version", apis.VersionPHP)
	r.GET("/versionjm", apis.VersionPHP)
	r.GET("/metrics", gin.BasicAuth(gin.Accounts{"fgadmin": "DHBOXlrZc71fOR8i"}), ginprom.PromHandler(promhttp.Handler()))

	a := r.Group("/h6hb7860q2")
	{
		a.GET("*", apis.Admin)
	}
	r.NoRoute(func(c *gin.Context) {
		c.Writer.WriteString("404")
		c.Status(http.StatusNotFound)
		c.Abort()
	})

	r.POST("/h6hb7860q2/api/login", apis.Login)
	admin := r.Group("/h6hb7860q2/api", middleware.JWTAuth())
	{
		admin.GET("menus", apis.Menus)
		admin.GET("versionInfo", apis.VersionInfo)
		admin.POST("/new/version", apis.NewVersion)
		admin.POST("/new/version/v2", apis.NewVersionV2)
		admin.POST("/new/domain", apis.NewDomain)
		admin.POST("/version", apis.Version)
		admin.GET("/get_templates", apis.TemplateList)
	}
}
