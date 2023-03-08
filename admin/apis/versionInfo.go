package apis

import (
	"github.com/LINQQ1212/common2/config/core"
	"github.com/LINQQ1212/common2/global"
	"github.com/LINQQ1212/common2/models"
	"github.com/LINQQ1212/common2/response"
	"github.com/gin-gonic/gin"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
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

func Version(c *gin.Context) {
	v := c.Query("v")
	version, ok := global.Versions.Get(v)
	if !ok {
		response.FailWithMessage("版本不存在", c)
		return
	}
	opt := core.VersionOption{}
	if err := c.BindJSON(&opt); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	version.Info.UseG = opt.UseG
	version.Info.UseY = opt.UseY
	version.Info.UseB = opt.UseB
	version.Info.UseYT = opt.UseYT
	version.Info.List = opt.List
	version.Info.Article = opt.Article
	version.Info.Option = opt.Option
	version.Info.CategoryLink = opt.CategoryLink
	version.Info.ProductLink = opt.ProductLink
	version.Info.UseBigSitemap = opt.UseBigSitemap
	version.Info.Category = opt.Category
	version.Info.RandTemp = opt.RandTemp
	version.Info.Paging = opt.Paging

	if version.Info.BigSitemap == nil {
		version.Info.BigSitemap = &models.Sitemap{}
	}
	version.Info.BigSitemap.Size = opt.BigSitemap.Size
	version.Info.BigSitemap.Option = opt.BigSitemap.Option
	if version.Info.BigSitemap.Size < 5 {
		version.Info.BigSitemap.Size = 5
	}

	if version.Info.SubSitemap == nil {
		version.Info.SubSitemap = &models.Sitemap{}
	}
	version.Info.SubSitemap.Size = opt.SubSitemap.Size
	version.Info.SubSitemap.Option = opt.SubSitemap.Option
	if version.Info.SubSitemap.Size < 500 {
		version.Info.SubSitemap.Size = 500
	}

	if version.Info.GoogleImg == nil {
		version.Info.GoogleImg = &models.External{}
	}
	version.Info.GoogleImg.Size = opt.GoogleImgs.Size
	version.Info.GoogleImg.Option = opt.GoogleImgs.Option
	version.Info.GoogleImg.GroupSize = opt.GoogleImgs.GroupSize
	if version.Info.GoogleImg.Size < 5 {
		version.Info.GoogleImg.Size = 5
	}

	vb, err := proto.Marshal(version.Info)
	if err != nil {
		global.LOG.Error("update vinfo", zap.Error(err))
		response.FailWithMessage("错误："+err.Error(), c)
	}
	err = version.DB.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket(models.BInfo).Put(models.BInfo, vb)
	})
	if err != nil {
		global.LOG.Error("update vinfo 2", zap.Error(err))
		response.FailWithMessage("错误："+err.Error(), c)
	}
	response.Ok(c)

}
